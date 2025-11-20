package api

type ChatRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}
