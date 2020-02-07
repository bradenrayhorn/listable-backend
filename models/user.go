package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/bradenrayhorn/listable-backend/utils"
)

type User struct {
	AutoId
	Name     string `json:"name"`
	Password string `json:"-"`
	ModelTimestamps
}

func (u User) TableName() string {
	return "users"
}

func (u User) MakeApiToken() string {
	token := ApiToken{}
	token.Generate()
	token.UserID = u.ID
	db.GetDB().Create(&token)
	return token.Token
}

func GetUser(username string) (User, error) {
	user := User{}
	count := 0
	err := db.GetDB().DB.Get(&count, "SELECT count(*) FROM users WHERE name = ?", username)
	if err != nil {
		return user, err
	}
	if count < 1 {
		return user, utils.ApiError{Code: 401, Reason: "invalid username / password"}
	}
	err = db.GetDB().DB.Get(&user, "SELECT * FROM users WHERE name = ?", username)
	return user, err
}
