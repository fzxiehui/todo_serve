package auth_service

import (
	"errors"

	"github.com/fzxiehui/todo_serve/internal/dal/model"
	"github.com/fzxiehui/todo_serve/internal/dal/query"
	"github.com/fzxiehui/todo_serve/internal/types"
	"github.com/fzxiehui/todo_serve/tools"
)

type Auth struct {
	Username string
	Nickname string
	Password string
}

func (a *Auth) Login() (*types.LoginResponse, error) {
	u := query.User
	user, err := u.Where(u.Username.Eq(a.Username)).First()
	if err != nil {
		return nil, err
	}

	if err := tools.CompareHash(user.Password, a.Password); err != nil {
		return nil, errors.New("password error")
	}

	token, err := tools.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Username: user.Username,
		Nickname: user.Nickname,
		Token:    token,
	}, nil
}

func (a *Auth) Register() (*types.RegisterResponse, error) {

	user := model.User{
		Username: a.Username,
		Nickname: a.Nickname,
		Password: tools.HashPassword(a.Password),
	}

	u := query.User
	err := u.Create(&user)
	if err != nil {
		return nil, err
	}

	token, err := tools.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		Username: user.Username,
		Nickname: user.Nickname,
		Token:    token,
	}, nil
}
