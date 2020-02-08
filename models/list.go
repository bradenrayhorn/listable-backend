package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
	"github.com/go-jet/jet/mysql"
)

type List struct {
	ID      uint32 `json:"id" alias:"lists.id"`
	GroupID uint32 `json:"-" alias:"lists.group_id"`
	Name    string `json:"name" alias:"lists.name"`

	ListItems []struct {
		ID      uint32 `json:"id"`
		Content string `json:"content"`
		Checked bool   `json:"checked"`
	} `json:"items" alias:"list_items.*"`
}

func GetAllListsForUser(userId uint32) ([]List, error) {
	var lists []List

	err := Lists.SELECT(
		Lists.ID,
		Lists.GroupID,
		Lists.Name,
		ListItems.ID,
		ListItems.Checked,
		ListItems.Content,
	).FROM(
		Lists.
			LEFT_JOIN(ListItems, ListItems.ListID.EQ(Lists.ID)).
			INNER_JOIN(Groups, Groups.ID.EQ(Lists.GroupID)).
			INNER_JOIN(GroupsUsers, GroupsUsers.GroupID.EQ(Groups.ID)).
			INNER_JOIN(Users, Users.ID.EQ(GroupsUsers.UserID))).
		WHERE(Users.ID.EQ(mysql.Int(int64(userId)))).Query(db.GetDB().DB, &lists)

	return lists, err
}
