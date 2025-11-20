package core

import "context"

type LLMProvider interface {
	Chat(ctx context.Context, prompt string) (string, error)
	StreamChat(ctx context.Context, prompt string, msgChan chan<- string) error
}
