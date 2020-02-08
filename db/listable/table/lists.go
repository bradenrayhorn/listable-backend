//
// Code generated by go-jet DO NOT EDIT.
// Generated at Saturday, 08-Feb-20 14:39:18 CST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/mysql"
)

var Lists = newListsTable()

type ListsTable struct {
	mysql.Table

	//Columns
	ID        mysql.ColumnInteger
	GroupID   mysql.ColumnInteger
	Name      mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new ListsTable with assigned alias
func (a *ListsTable) AS(alias string) *ListsTable {
	aliasTable := newListsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newListsTable() *ListsTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		GroupIDColumn   = mysql.IntegerColumn("group_id")
		NameColumn      = mysql.StringColumn("name")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
	)

	return &ListsTable{
		Table: mysql.NewTable("listable", "lists", IDColumn, GroupIDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn),

		//Columns
		ID:        IDColumn,
		GroupID:   GroupIDColumn,
		Name:      NameColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     mysql.ColumnList{IDColumn, GroupIDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn},
		MutableColumns: mysql.ColumnList{GroupIDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn},
	}
}
