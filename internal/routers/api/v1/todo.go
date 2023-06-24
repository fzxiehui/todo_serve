package v1

import (
	"net/http"
	"strconv"

	"github.com/fzxiehui/todo_serve/internal/types"
	"github.com/fzxiehui/todo_serve/services/todo_service"
	"github.com/gin-gonic/gin"
)

func GetTodoList(c *gin.Context) {

}

func GetTodo(c *gin.Context) {
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userid not found"})
		return
	}
	todoid := c.Param("id")
	todoidUint, err := strconv.ParseUint(todoid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := todo_service.Todo{
		UserId: userid.(uint),
		ID:     uint(todoidUint),
	}
	t, err := todo.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)

}

func CreateTodo(c *gin.Context) {
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userid not found"})
		return
	}

	var req types.CreateTodoRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := todo_service.Todo{
		UserId:  userid.(uint),
		Content: req.Content,
		Done:    req.Done,
	}

	t, err := todo.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}

func UpdateTodo(c *gin.Context) {
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userid not found"})
		return
	}

	todoid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req types.UpdateTodoRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := todo_service.Todo{
		UserId: userid.(uint),
		ID:     uint(todoid),
		Done:   req.Done,
	}

	t, err := todo.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}
