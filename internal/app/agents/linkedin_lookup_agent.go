package agents

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	agentTools "github.com/dragaoazuljr/ice-breaker-go/internal/app/tools"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/tools"
)

func Lookup(name string) string {
	model := os.Getenv("MODEL")

	llm, err := ollama.New(ollama.WithModel(model))

	if err != nil {
		log.Fatalf("Failed to load model: %v", err)
	}

	var template = fmt.Sprintf(`
    given the full name %s I want you to get it me a link to their LinkedIn profile page.
    Your answer should be a valid LinkedIn profile URL. 
		Answer only with the URL, do not include any other information.
    If the url has a different domain than www.linkedin.com, replace it with www, like br.linkedin.com will be www.linkedin.com 
	`, name)

	var tools = []tools.Tool{
		agentTools.GoogleSearch{},
	}

	executor, err := func() (*agents.Executor, error) {
		var opts []agents.Option = []agents.Option{agents.WithMaxIterations(3)}
		var agent agents.Agent
		switch agents.ZeroShotReactDescription {
		case agents.ZeroShotReactDescription:
			agent = agents.NewOneShotAgent(llms.Model(llm), tools, opts...)
		case agents.ConversationalReactDescription:
			agent = agents.NewConversationalAgent(llms.Model(llm), tools, opts...)
		default:
			return &agents.Executor{}, agents.ErrUnknownAgentType
		}
		return agents.NewExecutor(agent, tools, opts...), nil
	}()

	if err != nil {
		log.Fatalf("Failed to initialize agent: %v", err)
	}

	answer, err := chains.Run(context.Background(), executor, template)
	if err != nil {
		log.Fatalf("Failed to run agent: %v", err)
	}

	// Remove < from the start and > at the end from the answer
	if model == "mistral" {
		trim := strings.TrimSpace(answer)
		result := trim[1 : len(trim)-1]

		return result
	} else {
		return answer
	}
}
