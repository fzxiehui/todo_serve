package tools

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fzxiehui/todo_serve/config"
)

var jwtSecret = []byte(config.Config().GetString("jwt_secret"))

type Claims struct {
	UserID uint `json:"userid"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userid uint) (string, error) {

	nowTime := time.Now()

	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_serve",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
