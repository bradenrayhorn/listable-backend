package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/bradenrayhorn/listable-backend/utils"
	"math/rand"
	"time"
)

type ApiToken struct {
	UserID uint
	Token  string
	ModelTimestamps
}

func (token ApiToken) TableName() string {
	return "api_tokens"
}

func (token *ApiToken) Generate() {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, 64)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	token.Token = string(b)
}

func FindTokenUser(token string) (User, error) {
	user := User{}
	count := 0
	err := db.GetDB().DB.Get(&count, "SELECT count(*) FROM api_tokens WHERE token = ?", token)
	if err != nil {
		return user, err
	}
	if count < 1 {
		return user, utils.ApiError{Code: 401, Reason: "invalid token"}
	}
	err = db.GetDB().DB.Get(&user, "SELECT users.* FROM users join api_tokens on users.id = api_tokens.user_id WHERE token = ?", token)
	return user, err
}
