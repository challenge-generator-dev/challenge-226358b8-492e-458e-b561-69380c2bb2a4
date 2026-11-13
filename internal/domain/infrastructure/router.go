package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/application"
)

func SetupRouter(router *gin.Engine, taskService *application.TaskService) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("/", createTaskHandler(taskService))
		taskGroup.GET("/", listTasksHandler(taskService))
		taskGroup.GET("/:id", getTaskHandler(taskService))
		taskGroup.PUT("/:id", updateTaskHandler(taskService))
	}
}

func createTaskHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title       string `json:"title" binding:"required"`
			Description string `json:"description" binding:"required"`
			Status      string `json:"status" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		status := domain.TaskStatus(req.Status)
		task, err := taskService.CreateTask(c.Request.Context(), req.Title, req.Description, status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

func listTasksHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := taskService.ListTasks(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func getTaskHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		task, err := taskService.GetTask(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	}
}

func updateTaskHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var req struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Status      string `json:"status"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		status := domain.TaskStatus(req.Status)
		task, err := taskService.UpdateTask(c.Request.Context(), id, req.Title, req.Description, status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}