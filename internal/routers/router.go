package routers

import (
	"net/http"

	"github.com/fzxiehui/todo_serve/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/*
	 * Public
	 */
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	/*
	 * auth
	 */
	auth := router.Group("/auth")
	{
		auth.POST("/login", api.Login)
		auth.POST("/register", api.Register)
	}

	/*
	 * api v1
	 */

	return router
}
