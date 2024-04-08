package app

import (
	"fmt"

	agentLinkedinProfile "github.com/dragaoazuljr/ice-breaker-go/internal/app/agents"
	"github.com/dragaoazuljr/ice-breaker-go/internal/app/scrapper"
)

func IceBreaker() {
	var name = "Danillo Moraes"

	linkedInProfileUrl := agentLinkedinProfile.Lookup(name);
	linkedInData, error := scrapper.ScrapeLinkedInProfileData(linkedInProfileUrl);
	
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(linkedInData)
}
