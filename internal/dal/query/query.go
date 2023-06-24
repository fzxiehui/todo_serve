package query

import (
	"sync"

	"github.com/fzxiehui/todo_serve/internal/dal"
)

var once sync.Once

func init() {
	once.Do(func() {
		SetQuery()
	})
}

func SetQuery() {
	db := dal.GetDB()
	SetDefault(db)
}
