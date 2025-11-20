package main

import (
	"context"
	"fmt"
	"net/http"
	"neurogate/internal/api"
	"neurogate/internal/config"
	"neurogate/internal/infra/mock"
	"neurogate/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()

	logger.InitLogger(cfg.Server.Mode)
	defer logger.Log.Sync() // 确保缓冲日志写入
	logger.Log.Info("Starting Neurogate...",
		zap.String("version", cfg.App.Version),
		zap.String("env", cfg.Server.Mode),
	)

	llmClient := mock.NewMockClient()

	r := api.NewRouter(cfg, llmClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Listen: %s\n", zap.Error(err))
		}
	}()
	logger.Log.Info("Server started", zap.Int("port", cfg.Server.Port))

	quit := make(chan os.Signal, 1)
	// 注册监听，但系统向程序发送信号好，截取并进行后续操作
	// 阻塞监听是否有SIGINT或者SIGTERM系统信号传来
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server forced to shutdown: ", zap.Error(err))
	}

	logger.Log.Info("Server exiting")
}
