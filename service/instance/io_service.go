package instance

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	maxImportResponseBytes = 2 << 20
	maxImportRedirects     = 10
)

type lookupIPFunc func(context.Context, string) ([]net.IPAddr, error)

type IOService struct {
	client   *http.Client
	lookupIP lookupIPFunc
}

func NewIOService() *IOService {
	resolver := net.DefaultResolver
	lookupIP := resolver.LookupIPAddr

	dialer := &net.Dialer{Timeout: 5 * time.Second, KeepAlive: 30 * time.Second}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.Proxy = nil
	transport.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		host, port, err := net.SplitHostPort(address)
		if err != nil {
			return nil, fmt.Errorf("invalid destination address: %w", err)
		}
		addresses, err := lookupIP(ctx, host)
		if err != nil {
			return nil, fmt.Errorf("resolve destination: %w", err)
		}
		if err := validatePublicAddresses(addresses); err != nil {
			return nil, err
		}

		// Dial a validated address directly so a second DNS lookup cannot change
		// the destination between validation and connection (DNS rebinding).
		return dialer.DialContext(ctx, network, net.JoinHostPort(addresses[0].IP.String(), port))
	}
	transport.ResponseHeaderTimeout = 10 * time.Second
	transport.TLSHandshakeTimeout = 10 * time.Second
	transport.MaxIdleConnsPerHost = 2

	service := &IOService{lookupIP: lookupIP}
	service.client = &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= maxImportRedirects {
				return errors.New("too many redirects")
			}
			return service.validateURL(req.Context(), req.URL)
		},
	}
	return service
}

func (ioService IOService) ImportFromURL(ctx context.Context, rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parse URL: %w", err)
	}
	if err := ioService.validateURL(ctx, parsedURL); err != nil {
		return "", err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	response, err := ioService.client.Do(request)
	if err != nil {
		return "", fmt.Errorf("fetch URL: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %s", response.Status)
	}

	body, err := io.ReadAll(io.LimitReader(response.Body, maxImportResponseBytes+1))
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}
	if len(body) > maxImportResponseBytes {
		return "", fmt.Errorf("response exceeds %d bytes", maxImportResponseBytes)
	}
	document, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	return extractArticleText(document.Find("body").Children()), nil
}

func (ioService IOService) validateURL(ctx context.Context, parsedURL *url.URL) error {
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("only http and https URLs are allowed")
	}
	if parsedURL.Hostname() == "" {
		return errors.New("URL must include a host")
	}
	if parsedURL.User != nil {
		return errors.New("URL credentials are not allowed")
	}

	addresses, err := ioService.lookupIP(ctx, parsedURL.Hostname())
	if err != nil {
		return fmt.Errorf("resolve URL host: %w", err)
	}
	return validatePublicAddresses(addresses)
}

func validatePublicAddresses(addresses []net.IPAddr) error {
	if len(addresses) == 0 {
		return errors.New("URL host has no IP address")
	}
	for _, address := range addresses {
		ip := address.IP
		if ip == nil || ip.IsUnspecified() || ip.IsLoopback() || ip.IsPrivate() ||
			ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsMulticast() {
			return fmt.Errorf("URL host resolves to a non-public IP address: %s", ip)
		}
	}
	return nil
}

const articleNewLine = "\r\n"

func extractArticleText(children *goquery.Selection) string {
	if children == nil {
		return ""
	}
	const tags = "p, h1, h2, h3, h4, h5, h6, ul, ol, pre, blockquote"
	var result strings.Builder
	children.Each(func(_ int, selection *goquery.Selection) {
		if selection.Is(tags) {
			text := strings.TrimSpace(selection.Text())
			if text != "" {
				result.WriteString(text)
				result.WriteString(articleNewLine)
			}
			return
		}
		result.WriteString(extractArticleText(selection.Children()))
	})
	return result.String()
}
