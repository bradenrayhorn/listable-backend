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

var SchemaMigrations = newSchemaMigrationsTable()

type SchemaMigrationsTable struct {
	mysql.Table

	//Columns
	Version mysql.ColumnInteger
	Dirty   mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new SchemaMigrationsTable with assigned alias
func (a *SchemaMigrationsTable) AS(alias string) *SchemaMigrationsTable {
	aliasTable := newSchemaMigrationsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newSchemaMigrationsTable() *SchemaMigrationsTable {
	var (
		VersionColumn = mysql.IntegerColumn("version")
		DirtyColumn   = mysql.BoolColumn("dirty")
	)

	return &SchemaMigrationsTable{
		Table: mysql.NewTable("listable", "schema_migrations", VersionColumn, DirtyColumn),

		//Columns
		Version: VersionColumn,
		Dirty:   DirtyColumn,

		AllColumns:     mysql.ColumnList{VersionColumn, DirtyColumn},
		MutableColumns: mysql.ColumnList{DirtyColumn},
	}
}
