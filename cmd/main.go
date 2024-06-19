package main

import (
	"context"
	"fmt"
	"github.com/riyadennis/instructor/internal"
	"github.com/riyadennis/instructor/llm"
	"go.uber.org/zap"
)

var text = `Jane was travelling to London from her home in New York on a rainy day. She was 25 years old.`

func main() {
	ctx := context.Background()
	config, err := internal.NewConfig()
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	openAI := llm.NewOpenAIClient(logger, config.LlmApiKey, config.APIMaxRetries)
	person, err := openAI.ExtractPersonalInformation(ctx, text)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`
Name: %s
Age:  %d
Location: %s
`, person.Name, person.Age, person.Location)
}
