package http

import (
	"github.com/gin-gonic/gin"
	"github.com/thapakazi/go-hex-arch/internal/adapter/config"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()
	userHandler := NewUserHandler()
	routes := router.Group("/api")
	{
		user := routes.Group("/user")
		{
			user.GET("/:id", userHandler.GetUser)
			user.POST("", userHandler.CreateUser)
			user.DELETE("/:id", userHandler.DeleteUser)
			user.PUT("/:id", userHandler.UpdateUser)
			user.GET("", userHandler.GetAllUsers)
		}
	}

	return &Router{
		router,
	}
}

func (r *Router) Serve() error {

	return r.Run(":" + config.Environment.ServerPort)
}
