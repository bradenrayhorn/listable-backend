package controllers

import (
  "github.com/bradenrayhorn/listable-backend/models"
  "github.com/bradenrayhorn/listable-backend/utils"
  "github.com/gorilla/context"
  "net/http"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {
  
  user := context.Get(r, AuthUser).(models.User)
  
  utils.JsonResponse(w, "welcome "+user.Name)
}
