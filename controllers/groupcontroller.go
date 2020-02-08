package controllers

import (
	"github.com/bradenrayhorn/listable-backend/models"
	"github.com/bradenrayhorn/listable-backend/utils"
	"github.com/gorilla/context"
	"net/http"
)

func GetAllGroups(w http.ResponseWriter, r *http.Request) {

	user := context.Get(r, AuthUser).(models.User)
	groups, err := models.GetAllGroupsForUser(user.ID)

	if err != nil {
		utils.JsonError(err.Error(), w, http.StatusInternalServerError)
		return
	}

	utils.JsonSuccess(groups, w)
}

type CreateGroupRequest struct {
	Name string `validate:"required;unique:groups,name"`
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	rawValues, apiError := utils.ValidateRequest(CreateGroupRequest{}, r, w)
	if apiError != nil {
		utils.JsonError(apiError.Error(), w, apiError.Code)
		return
	}
	values := rawValues.(*CreateGroupRequest)

	err := models.CreateGroup(values.Name, context.Get(r, AuthUser).(models.User).ID)
	utils.CheckInternalError(w, err)
}
