package handler

import (
	"Brachium/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		checklist := api.Group("/checklists")
		{
			checklist.POST("/", h.createCheckList)
			checklist.GET("/", h.getAllCheckLists)
			checklist.GET("/:id", h.getCheckListById)
			checklist.PUT("/:id", h.updateCheckList)
			checklist.DELETE("/:id", h.deleteCheckList)

			tasks := checklist.Group(":id/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.GET("/", h.getAllTasks)
				tasks.GET("/:task_id", h.getTaskById)
				tasks.PUT("/:task_id", h.updateTask)
				tasks.DELETE("/:task_id", h.deleteTask)
			}
		}

	}
	return router
}
