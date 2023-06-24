package routers

import (
	"net/http"

	"github.com/fzxiehui/todo_serve/internal/middleware/jwt"
	"github.com/fzxiehui/todo_serve/internal/routers/api"
	v1 "github.com/fzxiehui/todo_serve/internal/routers/api/v1"
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
	apiv1 := router.Group("/v1")
	apiv1.Use(jwt.JWT())
	{
		/* todo */
		// apiv1.GET("/todo", v1.GetTodoList)
		apiv1.GET("/todo/:id", v1.GetTodo)
		apiv1.POST("/todo", v1.CreateTodo)
		apiv1.PATCH("/todo/:id", v1.UpdateTodo)
		apiv1.POST("/todoquery", v1.QueryTodo)
	}

	return router
}
