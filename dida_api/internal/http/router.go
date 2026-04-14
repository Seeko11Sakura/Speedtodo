package http

import (
	"dida_api/internal/focus"
	"dida_api/internal/list"
	"dida_api/internal/tag"
	"dida_api/internal/task"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	taskHandler := task.NewHandler()
	listHandler := list.NewHandler()
	tagHandler := tag.NewHandler()
	focusHandler := focus.NewHandler()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Tasks
	router.GET("/tasks", taskHandler.List)
	router.POST("/tasks", taskHandler.Create)
	router.GET("/tasks/:id", taskHandler.Get)
	router.PATCH("/tasks/:id", taskHandler.Update)
	router.POST("/tasks/:id/complete", taskHandler.Complete)
	router.POST("/tasks/:id/reopen", taskHandler.Reopen)
	router.DELETE("/tasks/:id", taskHandler.Delete)

	// Lists
	router.GET("/lists", listHandler.List)
	router.POST("/lists", listHandler.Create)
	router.PATCH("/lists/:id", listHandler.Update)

	// Tags
	router.GET("/tags", tagHandler.List)
	router.POST("/tags", tagHandler.Create)

	// Focus sessions
	router.GET("/focus-sessions", focusHandler.List)
	router.POST("/focus-sessions", focusHandler.Create)
	router.POST("/focus-sessions/:id/complete", focusHandler.Complete)
	router.POST("/focus-sessions/:id/cancel", focusHandler.Cancel)

	return router
}
