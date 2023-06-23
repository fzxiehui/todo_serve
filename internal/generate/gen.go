package main

import (
	"github.com/fzxiehui/todo_serve/internal/dal"
	"github.com/fzxiehui/todo_serve/internal/dal/model"
	"gorm.io/gen"
)

func main() {
	GenerateDal()
}

func GenerateDal() {
	config := gen.Config{
		OutPath:       "internal/dal/query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	}

	g := gen.NewGenerator(config)
	db := dal.GetDB()
	g.UseDB(db)

	g.ApplyBasic(model.User{})

	g.ApplyInterface(func() {}, model.User{})

	g.Execute()
}
