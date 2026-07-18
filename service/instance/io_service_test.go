package instance

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func publicLookup(_ context.Context, _ string) ([]net.IPAddr, error) {
	return []net.IPAddr{{IP: net.ParseIP("93.184.216.34")}}, nil
}

func TestImportFromURLRejectsUnsafeDestinations(t *testing.T) {
	tests := []struct {
		name      string
		rawURL    string
		addresses []net.IPAddr
	}{
		{name: "file scheme", rawURL: "file:///etc/passwd", addresses: nil},
		{name: "loopback", rawURL: "http://localhost/data", addresses: []net.IPAddr{{IP: net.ParseIP("127.0.0.1")}}},
		{name: "private", rawURL: "http://internal/data", addresses: []net.IPAddr{{IP: net.ParseIP("10.0.0.1")}}},
		{name: "link local", rawURL: "http://metadata/data", addresses: []net.IPAddr{{IP: net.ParseIP("169.254.169.254")}}},
		{name: "mixed DNS answers", rawURL: "https://example.test", addresses: []net.IPAddr{{IP: net.ParseIP("93.184.216.34")}, {IP: net.ParseIP("192.168.1.1")}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := IOService{
				client: &http.Client{Transport: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
					t.Fatal("unsafe URL reached HTTP transport")
					return nil, nil
				})},
				lookupIP: func(_ context.Context, _ string) ([]net.IPAddr, error) { return test.addresses, nil },
			}
			if _, err := service.ImportFromURL(context.Background(), test.rawURL); err == nil {
				t.Fatal("expected unsafe URL to be rejected")
			}
		})
	}
}

func TestImportFromURLExtractsArticleText(t *testing.T) {
	service := IOService{
		lookupIP: publicLookup,
		client: &http.Client{Transport: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(strings.NewReader("<html><body><h1>Title</h1><div><p>Paragraph</p></div></body></html>")),
			}, nil
		})},
	}

	content, err := service.ImportFromURL(context.Background(), "https://example.test/article")
	if err != nil {
		t.Fatalf("ImportFromURL returned error: %v", err)
	}
	if want := "Title\r\nParagraph\r\n"; content != want {
		t.Fatalf("content = %q, want %q", content, want)
	}
}

func TestImportFromURLLimitsResponseSize(t *testing.T) {
	service := IOService{
		lookupIP: publicLookup,
		client: &http.Client{Transport: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(strings.NewReader(strings.Repeat("x", maxImportResponseBytes+1))),
			}, nil
		})},
	}

	if _, err := service.ImportFromURL(context.Background(), "https://example.test/large"); err == nil {
		t.Fatal("expected oversized response to be rejected")
	}
}

func TestRedirectPolicyRejectsPrivateDestination(t *testing.T) {
	service := NewIOService()
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1/secret", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := service.client.CheckRedirect(request, nil); err == nil {
		t.Fatal("expected redirect to private destination to be rejected")
	}
}
