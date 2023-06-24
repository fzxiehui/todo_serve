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

	/*
	 * 生成单表
	 */
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo{})

	/*
	 * 生成基本类型安全的 DAO API
	 */
	g.ApplyBasic(model.User{})
	g.ApplyBasic(model.Todo{})

	/*
	 * 生成自定义的 DAO API
	 */

	// g.ApplyInterface(func() {}, model.User{})

	g.Execute()
}
