package controllers

import (
  "github.com/bradenrayhorn/listable-backend/models"
  "github.com/bradenrayhorn/listable-backend/utils"
  "golang.org/x/crypto/bcrypt"
  "net/http"
)

type LoginRequest struct {
  Username string `validate:"required"`
  Password string `validate:"required"`
}

func Login(w http.ResponseWriter, r *http.Request) {
  rawValues, apiError := utils.ValidateRequest(LoginRequest{}, r, w)
  if apiError != nil {
    utils.JsonError(apiError.Error(), w, apiError.Code)
    return
  }
  values := rawValues.(*LoginRequest)
  
  user, err := models.GetUser(values.Username)
  
  if err != nil {
    utils.JsonError(err.Error(), w, 500)
    return
  }
  
  success := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(values.Password))
  
  if success != nil {
    utils.JsonError("invalid username / password", w, 401)
  } else {
    token := user.MakeApiToken()
    utils.JsonSuccess(map[string]string{"token": token}, w)
  }
}
