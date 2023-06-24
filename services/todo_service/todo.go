package todo_service

import (
	"time"

	"github.com/fzxiehui/todo_serve/internal/dal/model"
	"github.com/fzxiehui/todo_serve/internal/dal/query"
	"github.com/fzxiehui/todo_serve/internal/types"
)

type Todo struct {
	ID      uint
	Date    string
	Content string
	Done    bool
	UserId  uint
}

func (t *Todo) Create() (*types.CreateTodoResponse, error) {

	todo := model.Todo{
		Date:    time.Now().Format("2006-01-02"),
		Content: t.Content,
		Done:    t.Done,
		UserId:  t.UserId,
	}

	qt := query.Todo
	err := qt.Create(&todo)
	if err != nil {
		return nil, err
	}

	return &types.CreateTodoResponse{
		ID:      todo.ID,
		Date:    todo.Date,
		Content: todo.Content,
		Done:    todo.Done,
	}, nil
}

func (t *Todo) Get() (*types.GetTodoResponse, error) {

	qt := query.Todo
	todo, err := qt.Where(qt.ID.Eq(t.ID), qt.UserId.Eq(t.UserId)).First()
	if err != nil {
		return nil, err
	}

	return &types.GetTodoResponse{
		ID:      todo.ID,
		Date:    todo.Date,
		Content: todo.Content,
		Done:    todo.Done,
	}, nil

}

func (t *Todo) Update() (*types.UpdateTodoResponse, error) {

	// qt := query.Todo
	// user, err := qt.Where(qt.ID.Eq(t.ID),
	// 	qt.UserId.Eq(t.UserId)).Update(qt.Done, t.Done)
	// if err != nil {
	// 	return nil, err
	// }
	qt := query.Todo
	user, err := qt.Where(qt.ID.Eq(t.ID),
		qt.UserId.Eq(t.UserId)).First()

	if err != nil {
		return nil, err
	}

	user.Done = t.Done

	qt.Save(user)

	return &types.UpdateTodoResponse{
		ID:      user.ID,
		Date:    user.Date,
		Content: user.Content,
		Done:    user.Done,
	}, nil

}
