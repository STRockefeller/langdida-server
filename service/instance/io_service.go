package instance

import (
	"context"
	"io/ioutil"

	crawler "github.com/STRockefeller/article-crawler"
)

type IOService struct{}

func NewIOService() *IOService {
	return &IOService{}
}

func (io IOService) ImportFromURL(ctx context.Context, url string) (string, error) {
	return crawler.Crawl(url)
}

func (io IOService) ImportFromFile(ctx context.Context, filePath string) (string, error) {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
