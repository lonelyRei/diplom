package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// Соль для хэширования пароля
	salt = "cskdcmksckldsmkcd3991&&&!/"

	// Время жизни токена
	tokenTTL = 12 * time.Hour

	// Ключ сигнатуры
	signedKey = "cmsdmcscsd"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
