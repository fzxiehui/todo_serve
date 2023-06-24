package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fzxiehui/todo_serve/log"
	"github.com/fzxiehui/todo_serve/tools"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		// log.Info("token: ", token)
		if token == "" {
			token, _ = c.Cookie("Authorization")
			// log.Debug("token from cookie: ", token)
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty"})
			log.Debug("token is empty")
			c.Abort()
			return
		} else {
			// parse token
			user, err := tools.ParseToken(token)
			c.Set("userid", user.UserID)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "token is expired"})
					c.Abort()
					return
				default:

					c.JSON(http.StatusUnauthorized, gin.H{
						"error": err.Error()})
					c.Abort()
					return
				}
			}

			referer := c.Request.Header.Get("Referer")
			// set_token, err := tools.GenerateToken(user.UserID)
			set_token, err := tools.GenerateToken(user.UserID)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error()})
				c.Abort()
				return
			}
			c.SetCookie("Authorization",
				set_token, 3600, "/", referer, false, false)
		}
		// if token is valid, call next handler
		c.Next()
	}
}
