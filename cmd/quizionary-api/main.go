package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhori/go-quizionary-api/internal/handler"
)

func main() {
	// Ginのデフォルトモードを設定
	gin.SetMode(gin.ReleaseMode)

	// Handler のインスタンスを作成
	h, err := handler.New("files/four_option")
	if err != nil {
		panic(fmt.Errorf("failed to create handler: %w", err))
	}

	// Routerの設定
	r := gin.Default()
	r.GET("/four-option-questions", h.GetFourOptionQuizzes)
	r.GET("/ok", h.OK)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// HTTPサーバの設定
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// 別ゴルーチンでサーバ起動
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("failed to start server: %w", err))
		}
	}()
	slog.Info("Server started", "port", port)

	// シグナル受信用チャネル
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	// コンテキストタイムアウト付きでシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	slog.Info("Server exited gracefully")
}
