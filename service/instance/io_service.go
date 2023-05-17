package instance

import (
	"context"

	crawler "github.com/STRockefeller/article-crawler"
)

type IOService struct{}

func NewIOService() *IOService {
	return &IOService{}
}

func (io IOService) ImportFromURL(ctx context.Context, url string) (string, error) {
	return crawler.Crawl(url)
}
