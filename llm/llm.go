package llm

import "context"

type ChatMessage struct {
	Model   string
	Role    string
	Content string
}
type LLM interface {
	ExtractPersonalInformation(ctx context.Context, content string) (*Person, error)
}
