package api

import (
	"net/http"

	"github.com/fzxiehui/todo_serve/internal/types"
	"github.com/fzxiehui/todo_serve/services/auth_service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var req types.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth := auth_service.Auth{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := auth.Login()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	referer := c.Request.Header.Get("Referer")
	c.SetCookie("Authorization", user.Token, 3600, "/", referer, false, false)

	c.JSON(http.StatusOK, user)
}

func Register(c *gin.Context) {

	var req types.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth := auth_service.Auth{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
	}

	user, err := auth.Register()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
