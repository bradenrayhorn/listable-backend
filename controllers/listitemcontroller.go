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

type SetListItemCheckedRequest struct {
	ListItemId int  `validate:"required;integer"`
	Checked    bool `validate:"required;boolean"`
}

type RemoveListItemRequest struct {
	ListItemId int `validate:"required;integer"`
}

type ClearListRequest struct {
	ListId int `validate:"required;integer"`
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

func SetListItemChecked(w http.ResponseWriter, r *http.Request) {
	rawValues, apiError := utils.ValidateRequest(SetListItemCheckedRequest{}, r, w)
	if apiError != nil {
		utils.JsonError(apiError.Error(), w, apiError.Code)
		return
	}
	values := rawValues.(*SetListItemCheckedRequest)

	user := context.Get(r, AuthUser).(models.User)

	hasListItem, err := models.UserHasListItem(user.ID, values.ListItemId)
	if err != nil {
		utils.JsonError(err.Error(), w, http.StatusInternalServerError)
		return
	}
	if !hasListItem {
		utils.JsonError("you do not have access to this list item", w, http.StatusForbidden)
		return
	}

	if err = models.SetListItemChecked(values.ListItemId, values.Checked); err != nil {
		utils.JsonError("failed to set checked state", w, http.StatusInternalServerError)
	}
}

func RemoveListItem(w http.ResponseWriter, r *http.Request) {
	rawValues, apiError := utils.ValidateRequest(RemoveListItemRequest{}, r, w)
	if apiError != nil {
		utils.JsonError(apiError.Error(), w, apiError.Code)
		return
	}
	values := rawValues.(*RemoveListItemRequest)

	user := context.Get(r, AuthUser).(models.User)

	hasListItem, err := models.UserHasListItem(user.ID, values.ListItemId)
	if err != nil {
		utils.JsonError(err.Error(), w, http.StatusInternalServerError)
		return
	}
	if !hasListItem {
		utils.JsonError("you do not have access to this list item", w, http.StatusForbidden)
		return
	}

	if err = models.RemoveListItem(values.ListItemId); err != nil {
		utils.JsonError("failed to remove list item", w, http.StatusInternalServerError)
	}
}

func ClearList(w http.ResponseWriter, r *http.Request) {
	rawValues, apiError := utils.ValidateRequest(ClearListRequest{}, r, w)
	if apiError != nil {
		utils.JsonError(apiError.Error(), w, apiError.Code)
		return
	}
	values := rawValues.(*ClearListRequest)

	user := context.Get(r, AuthUser).(models.User)

	hasListItem, err := models.UserHasList(user.ID, values.ListId)
	if err != nil {
		utils.JsonError(err.Error(), w, http.StatusInternalServerError)
		return
	}
	if !hasListItem {
		utils.JsonError("you do not have access to this list", w, http.StatusForbidden)
		return
	}

	if err = models.RemoveListItemsForList(values.ListId); err != nil {
		utils.JsonError("failed to remove list items", w, http.StatusInternalServerError)
	}
}
