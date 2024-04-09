package tools

import (
	"context"
	
	"github.com/tmc/langchaingo/tools"
	"github.com/dragaoazuljr/ice-breaker-go/internal/app/scrapper"
)

type GoogleSearchTwitter struct {
}

var _ tools.Tool = GoogleSearchTwitter{}

func (g GoogleSearchTwitter) Name() string {
	return "Crawl Google 4 Twitter profile pages"
}

func (g GoogleSearchTwitter) Description() string {
	return "Useful for when you need to find a link of Twitter profile for a person"
}

func (g GoogleSearchTwitter) Call(ctx context.Context, input string) (string, error) {
	return scrapper.GetFirstGoogleResultLink(input, "Twitter")
}

