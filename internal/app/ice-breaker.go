package app

import (
	"context"
	"fmt"
	"os"

	agents "github.com/dragaoazuljr/ice-breaker-go/internal/app/agents"
	"github.com/dragaoazuljr/ice-breaker-go/internal/app/scrapper"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/outputparser"
	"github.com/tmc/langchaingo/prompts"
)

func IceBreaker() {
	var name = "Danillo Moraes"

	linkedInProfileUrl := agents.LookupLinkedIn(name);
	linkedInData, error := scrapper.ScrapeLinkedInProfileData(linkedInProfileUrl);

	 if error != nil {
	 	fmt.Println(error)
	  	return
	 }

	twitterProfileUrl := agents.LookupTwitter(name);
	twitterData, error := scrapper.ScapeTwitterProfileData(twitterProfileUrl);
	
	if error != nil {
		fmt.Println(error)
		return
	}

	var summary_template = `
Use the following pieces of information about a person and answer the request.
----
LinkedIn Information: {{.linkedin_information}}
----
Twitter Information: {{.twitter_information}}
----

Request:
1. A short summary about the work experience of the person
2. A list of interesting related facts about him
3. A list of topics that you think they would be interested in discussing
4. A list of creative Ice Breakers to open a conversation with him

{{.formated_instructions}}
`

	var summary_prompt_template = prompts.NewPromptTemplate(
		summary_template,
		[]string{"linkedin_information", "twitter_information", "formated_instructions"},
	)

	model := os.Getenv("MODEL")
	llm, err:= ollama.New(ollama.WithModel(model))

	if err != nil {
		fmt.Println(err)
		return
	}

	responseSchemas :=[]outputparser.ResponseSchema{
		{ Name: "summary", Description: "Summary of the person" },
		{ Name: "facts", Description: "A list of interesting facts about the person" },
		{ Name: "topics", Description: "A list of topics that may interest for the person" },
		{ Name: "ice_breakers", Description: "A list of ice breakers to open a conversation with the person" },
	}

	structureParser := outputparser.NewStructured(responseSchemas)

	chain := chains.NewLLMChain(
		llm, 
		summary_prompt_template,
	)

	ctx:= context.Background()

	res, err := chain.Call(ctx, map[string]any{
		"linkedin_information": linkedInData,
		"twitter_information": twitterData,
		"formated_instructions": structureParser.GetFormatInstructions(),
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	response := res["text"].(string)

	formatedResponse, err := structureParser.Parse(response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(formatedResponse)
}
