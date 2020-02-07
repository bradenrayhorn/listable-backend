package controllers

import (
	"github.com/bradenrayhorn/listable-backend/models"
	"github.com/bradenrayhorn/listable-backend/utils"
	"github.com/gorilla/context"
	"net/http"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {

	user := context.Get(r, AuthUser).(models.User)

	lists, err := models.GetAllListsForUser(user.ID)

	if err != nil {
		utils.JsonError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	utils.JsonSuccess(lists, w)
}
