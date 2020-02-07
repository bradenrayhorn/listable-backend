package controllers

import (
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/bradenrayhorn/listable-backend/models"
	"github.com/bradenrayhorn/listable-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterRequest struct {
	Username string `validate:"required;unique:users,name"`
	Password string `validate:"required"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	rawValues, apiError := utils.ValidateRequest(RegisterRequest{}, r, w)
	if apiError != nil {
		utils.JsonError(apiError.Error(), w, apiError.Code)
		return
	}
	values := rawValues.(*RegisterRequest)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(values.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JsonError(err.Error(), w, 500)
		return
	}

	user := models.User{
		Name:     values.Username,
		Password: string(hashedPassword),
	}
	db.GetDB().Create(&user)
	user.MakeApiToken()
}
