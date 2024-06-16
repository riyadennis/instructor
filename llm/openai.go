package llm

import (
	"context"
	"github.com/instructor-ai/instructor-go/pkg/instructor"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

type OpenAIClient struct {
	APIKey   string
	MaxTries int
	Client   *instructor.InstructorOpenAI
	Logger   *zap.Logger
}

type Person struct {
	Name     string `json:"name"          jsonschema:"title=the name,description=The name of the person,example=joe,example=lucy"`
	Age      int    `json:"age,omitempty" jsonschema:"title=the age,description=The age of the person,example=25,example=67"`
	Location string `json:"location,omitempty" jsonschema:"location=the location,description=The location of the person,example=wimbledon,example=wimbledon"`
}

func NewOpenAIClient(logger *zap.Logger, key string, maxTries int) *OpenAIClient {
	return &OpenAIClient{
		APIKey:   key,
		MaxTries: maxTries,
		Logger:   logger,
		Client: instructor.FromOpenAI(
			openai.NewClient(key),
			instructor.WithMode(instructor.ModeJSON),
			instructor.WithMaxRetries(maxTries),
		),
	}
}

func (oai *OpenAIClient) ExtractPersonalInformation(ctx context.Context, content string) (*Person, error) {
	per := &Person{}
	_, err := oai.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Extract name, age and location:" + content,
				},
			},
		},
		&per,
	)
	if err != nil {
		return nil, err
	}
	return per, nil
}
