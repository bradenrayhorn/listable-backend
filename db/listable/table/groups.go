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

var Groups = newGroupsTable()

type GroupsTable struct {
	mysql.Table

	//Columns
	ID        mysql.ColumnInteger
	Name      mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new GroupsTable with assigned alias
func (a *GroupsTable) AS(alias string) *GroupsTable {
	aliasTable := newGroupsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newGroupsTable() *GroupsTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		NameColumn      = mysql.StringColumn("name")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
	)

	return &GroupsTable{
		Table: mysql.NewTable("listable", "groups", IDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn),

		//Columns
		ID:        IDColumn,
		Name:      NameColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     mysql.ColumnList{IDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn},
		MutableColumns: mysql.ColumnList{NameColumn, CreatedAtColumn, UpdatedAtColumn},
	}
}
