package controllers

import (
	"github.com/bradenrayhorn/listable-backend/models"
	"github.com/bradenrayhorn/listable-backend/utils"
	"github.com/gorilla/context"
	"net/http"
)

type AddListItemRequest struct {
	Content string `validate:"required"`
	List    int    `validate:"required;integer"`
}

func AddListItem(w http.ResponseWriter, r *http.Request) {
	rawValues, apiError := utils.ValidateRequest(AddListItemRequest{}, r, w)
	if apiError != nil {
		utils.JsonError(apiError.Error(), w, apiError.Code)
		return
	}
	values := rawValues.(*AddListItemRequest)

	user := context.Get(r, AuthUser).(models.User)

	hasList, err := models.UserHasList(user.ID, values.List)
	if err != nil {
		utils.JsonError(err.Error(), w, http.StatusInternalServerError)
		return
	}
	if !hasList {
		utils.JsonError("you do not have access to this list", w, http.StatusForbidden)
		return
	}

	if err = models.AddListItem(values.List, values.Content); err != nil {
		utils.JsonError("failed to add item to list", w, http.StatusInternalServerError)
	}
}
