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

var Users = newUsersTable()

type UsersTable struct {
	mysql.Table

	//Columns
	ID        mysql.ColumnInteger
	Name      mysql.ColumnString
	Password  mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new UsersTable with assigned alias
func (a *UsersTable) AS(alias string) *UsersTable {
	aliasTable := newUsersTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newUsersTable() *UsersTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		NameColumn      = mysql.StringColumn("name")
		PasswordColumn  = mysql.StringColumn("password")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
	)

	return &UsersTable{
		Table: mysql.NewTable("listable", "users", IDColumn, NameColumn, PasswordColumn, CreatedAtColumn, UpdatedAtColumn),

		//Columns
		ID:        IDColumn,
		Name:      NameColumn,
		Password:  PasswordColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     mysql.ColumnList{IDColumn, NameColumn, PasswordColumn, CreatedAtColumn, UpdatedAtColumn},
		MutableColumns: mysql.ColumnList{NameColumn, PasswordColumn, CreatedAtColumn, UpdatedAtColumn},
	}
}
