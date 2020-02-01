package controllers

import (
  "encoding/json"
  "github.com/bradenrayhorn/listable-backend/models"
  "net/http"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(models.GetAllLists())
}
