package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
	"github.com/go-jet/jet/mysql"
)

func AddListItem(listId int, content string) error {
	_, err := ListItems.
		INSERT(ListItems.ListID, ListItems.Content).
		VALUES(listId, content).
		Exec(db.GetDB().DB)

	return err
}

func SetListItemChecked(listItemId int, checked bool) error {
	_, err := ListItems.
		UPDATE(ListItems.Checked).
		SET(checked).
		WHERE(ListItems.ID.EQ(mysql.Int(int64(listItemId)))).
		Exec(db.GetDB().DB)

	return err
}

func RemoveListItem(listItemId int) error {
	_, err := ListItems.
		DELETE().
		WHERE(ListItems.ID.EQ(mysql.Int(int64(listItemId)))).
		Exec(db.GetDB().DB)

	return err
}

func RemoveListItemsForList(listId int) error {
	_, err := ListItems.
		DELETE().
		WHERE(ListItems.ID.IN(ListItems.INNER_JOIN(Lists, Lists.ID.EQ(ListItems.ListID)).SELECT(ListItems.ID).WHERE(Lists.ID.EQ(mysql.Int(int64(listId)))))).
		Exec(db.GetDB().DB)

	return err
}

func UserHasListItem(userId uint32, listItemId int) (bool, error) {
	result := struct {
		Count int
	}{0}
	err := ListItems.
		SELECT(mysql.COUNT(ListItems.ID).AS("count")).
		FROM(ListItems.
			INNER_JOIN(Lists, Lists.ID.EQ(ListItems.ListID)).
			INNER_JOIN(Groups, Groups.ID.EQ(Lists.GroupID)).
			INNER_JOIN(GroupsUsers, GroupsUsers.GroupID.EQ(Groups.ID)).
			INNER_JOIN(Users, Users.ID.EQ(GroupsUsers.UserID))).
		WHERE(Users.ID.EQ(mysql.Int(int64(userId)))).
		WHERE(ListItems.ID.EQ(mysql.Int(int64(listItemId)))).
		Query(db.GetDB().DB, &result)

	return result.Count > 0, err
}
