package api

import (
	"net/http"
	"neurogate/internal/core"
	"neurogate/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChatHandler struct {
	llm core.LLMProvider
}

func NewChatHandler(llm core.LLMProvider) *ChatHandler {
	return &ChatHandler{
		llm: llm,
	}
}

// Chat 普通对话接口的处理函数
// @Summary 普通对话
// @Description 接收 Prompt，返回完整回复
// @Accept json
// @Produce json
func (h *ChatHandler) Chat(c *gin.Context) {
	var req ChatRequest
	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Warn("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.Request.Context()传递了 HTTP 请求的上下文，如果用户断开，Context 会自动断开
	reply, err := h.llm.Chat(c.Request.Context(), req.Prompt)
	if err != nil {
		logger.Log.Error("LLM chat failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI service unavailable"})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{
		Reply: reply,
	})
}
