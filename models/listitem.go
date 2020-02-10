package models

import (
	"github.com/bradenrayhorn/listable-backend/db"
	. "github.com/bradenrayhorn/listable-backend/db/listable/table"
)

func AddListItem(listId int, content string) error {
	_, err := ListItems.
		INSERT(ListItems.ListID, ListItems.Content).
		VALUES(listId, content).
		Exec(db.GetDB().DB)

	return err
}
