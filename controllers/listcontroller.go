package controllers

import (
  "github.com/bradenrayhorn/listable-backend/models"
  "github.com/bradenrayhorn/listable-backend/utils"
  "net/http"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {
  utils.JsonResponse(w, models.GetAllLists())
}
