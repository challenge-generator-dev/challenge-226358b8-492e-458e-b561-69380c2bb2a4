package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure"
)

func main() {
	router := gin.Default()

	repo := infrastructure.NewInMemoryTaskRepository()
	taskService := application.NewTaskService(repo)

	infrastructure.SetupRouter(router, taskService)

	log.Println("Servidor iniciando en el puerto 8080")
	if err := router.Run(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}