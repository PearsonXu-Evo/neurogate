package mock

import (
	"context"
	"fmt"
	"neurogate/internal/core"
	"time"
)

type MockClient struct{}

var _ core.LLMProvider = (*MockClient)(nil)

func NewMockClient() core.LLMProvider {
	return &MockClient{}
}

func (c *MockClient) Chat(ctx context.Context, prompt string) (string, error) {
	return fmt.Sprintf("【Mock回复】我收到了你的问题：%s", prompt), nil
}

func (c *MockClient) StreamChat(ctx context.Context, prompt string, msgChan chan<- string) error {
	go func() {
		defer close(msgChan)
		response := []string{"你好", "，", "我是", "NeuroGate", "智能", "助手", "（", "Mock", "版", "）", "。"}
		for _, token := range response {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(100 * time.Millisecond)
				msgChan <- token
			}
		}
	}()
	return nil
}
