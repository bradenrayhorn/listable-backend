//
// Code generated by go-jet DO NOT EDIT.
// Generated at Saturday, 08-Feb-20 15:43:50 CST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/mysql"
)

var ListItems = newListItemsTable()

type ListItemsTable struct {
	mysql.Table

	//Columns
	ID        mysql.ColumnInteger
	ListID    mysql.ColumnInteger
	Content   mysql.ColumnString
	Checked   mysql.ColumnBool
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new ListItemsTable with assigned alias
func (a *ListItemsTable) AS(alias string) *ListItemsTable {
	aliasTable := newListItemsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newListItemsTable() *ListItemsTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		ListIDColumn    = mysql.IntegerColumn("list_id")
		ContentColumn   = mysql.StringColumn("content")
		CheckedColumn   = mysql.BoolColumn("checked")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
	)

	return &ListItemsTable{
		Table: mysql.NewTable("listable", "list_items", IDColumn, ListIDColumn, ContentColumn, CheckedColumn, CreatedAtColumn, UpdatedAtColumn),

		//Columns
		ID:        IDColumn,
		ListID:    ListIDColumn,
		Content:   ContentColumn,
		Checked:   CheckedColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     mysql.ColumnList{IDColumn, ListIDColumn, ContentColumn, CheckedColumn, CreatedAtColumn, UpdatedAtColumn},
		MutableColumns: mysql.ColumnList{ListIDColumn, ContentColumn, CheckedColumn, CreatedAtColumn, UpdatedAtColumn},
	}
}
