package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/bradenrayhorn/listable-backend/db/listable/model"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
	"github.com/bradenrayhorn/listable-backend/utils"
	"github.com/go-jet/jet/mysql"
	"math/rand"
	"time"
)

type ApiToken struct {
	model.APITokens
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

func (token ApiToken) Create() error {
	_, err := APITokens.INSERT(APITokens.UserID, APITokens.Token).MODEL(token).Exec(db.GetDB().DB)
	return err
}

func FindTokenUser(token string) (User, error) {
	var users []User
	err := Users.SELECT(Users.AllColumns).FROM(Users.
		INNER_JOIN(APITokens, APITokens.UserID.EQ(Users.ID))).
		WHERE(APITokens.Token.EQ(mysql.String(token))).Query(db.GetDB().DB, &users)
	if err != nil {
		return User{}, err
	}
	if len(users) != 1 {
		return User{}, utils.ApiError{Code: 401, Reason: "invalid token"}
	}
	return users[0], err
}
