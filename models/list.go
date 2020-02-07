package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
)

type List struct {
	AutoId
	Name    string `json:"name" db:"name"`
	GroupID uint   `json:"-" db:"group_id"`
	ModelTimestamps
}

func GetAllListsForUser(userId uint) ([]List, error) {
	var lists []List

	err := db.GetDB().DB.Select(&lists, "SELECT lists.* FROM lists join `groups` on lists.group_id = `groups`.id join groups_users gu on `groups`.id = gu.group_id join users u on gu.user_id = u.id where u.id = ?", userId)

	return lists, err
}
