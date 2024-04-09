package tools

import (
	"context"
	
	"github.com/tmc/langchaingo/tools"
	"github.com/dragaoazuljr/ice-breaker-go/internal/app/scrapper"
)

type GoogleSearchLinkedIn struct {
}

var _ tools.Tool = GoogleSearchLinkedIn{}

func (g GoogleSearchLinkedIn) Name() string {
	return "Crawl Google 4 LinkedIn profile pages"
}

func (g GoogleSearchLinkedIn) Description() string {
	return "Useful for when you need to find a link of LinkedIn profile for a person"
}

func (g GoogleSearchLinkedIn) Call(ctx context.Context, input string) (string, error) {
	return scrapper.GetFirstGoogleResultLink(input, "LinkedIn")
}
