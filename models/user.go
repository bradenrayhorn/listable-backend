package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/bradenrayhorn/listable-backend/db/listable/model"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
	"github.com/bradenrayhorn/listable-backend/utils"
	. "github.com/go-jet/jet/mysql"
)

type User struct {
	model.Users
}

func (u User) MakeApiToken() (string, error) {
	token := ApiToken{}
	token.Generate()
	token.UserID = u.ID
	err := token.Create()
	return token.Token, err
}

func (u *User) Create() error {
	res, err := Users.INSERT(Users.Name, Users.Password).MODEL(u).Exec(db.GetDB().DB)
	if err != nil {
		return utils.ApiError{Code: 500, Reason: err.Error()}
	}
	id, insertErr := res.LastInsertId()
	if insertErr != nil {
		return utils.ApiError{Code: 500, Reason: insertErr.Error()}
	}
	u.ID = uint32(id)
	return nil
}

func GetUser(username string) (User, error) {
	var users []User
	err := SELECT(Users.AllColumns).FROM(Users).WHERE(Users.Name.EQ(String(username))).Query(db.GetDB().DB, &users)
	if err != nil {
		return User{}, err
	}
	if len(users) != 1 {
		return User{}, utils.ApiError{Code: 401, Reason: "invalid username / password"}
	}
	return users[0], nil
}
