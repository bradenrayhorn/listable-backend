package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/bradenrayhorn/listable-backend/db/listable/model"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
	"github.com/go-jet/jet/mysql"
)

type List struct {
	model.Lists
}

func GetAllListsForUser(userId uint32) ([]List, error) {
	var lists []List

	err := Lists.SELECT(Lists.AllColumns).FROM(
		Lists.
			INNER_JOIN(Groups, Groups.ID.EQ(Lists.GroupID)).
			INNER_JOIN(GroupsUsers, GroupsUsers.GroupID.EQ(Groups.ID)).
			INNER_JOIN(Users, Users.ID.EQ(GroupsUsers.UserID))).
		WHERE(Users.ID.EQ(mysql.Int(int64(userId)))).Query(db.GetDB().DB, &lists)

	if len(lists) == 0 {
		return make([]List, 0), err
	}
	return lists, err
}
