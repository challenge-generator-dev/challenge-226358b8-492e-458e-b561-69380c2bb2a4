package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	taskRepo := infrastructure.NewInMemoryTaskRepository()

	taskService := application.NewTaskService(taskRepo)

	infrastructure.SetupRouter(router, taskService)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server starting on port %s", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server exited")
}