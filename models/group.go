package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
	"github.com/go-jet/jet/mysql"
)

type Group struct {
	ID   uint32 `json:"id" alias:"groups.id"`
	Name string `json:"name" alias:"groups.name"`
}

func GetAllGroupsForUser(userId uint32) ([]Group, error) {
	var groups []Group

	err := Groups.SELECT(Groups.AllColumns).FROM(
		Groups.
			INNER_JOIN(GroupsUsers, GroupsUsers.GroupID.EQ(Groups.ID))).
		WHERE(GroupsUsers.UserID.EQ(mysql.Int(int64(userId)))).Query(db.GetDB().DB, &groups)

	if len(groups) == 0 {
		return make([]Group, 0), err
	}
	return groups, err
}

func CreateGroup(groupName string, userId uint32) error {
	res, err := Groups.INSERT(Groups.Name).VALUES(groupName).Exec(db.GetDB().DB)
	if err != nil {
		return err
	}
	groupId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = GroupsUsers.INSERT(GroupsUsers.UserID, GroupsUsers.GroupID).VALUES(userId, groupId).Exec(db.GetDB().DB)
	return err
}
