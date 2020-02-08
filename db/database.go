package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"time"
)

type DB struct {
	DB *sqlx.DB
}

type HasTableName interface {
	TableName() string
}

var db *DB

func SetupDatabase() {
	dbCon, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql_username"),
		viper.GetString("mysql_password"),
		viper.GetString("mysql_host"),
		viper.GetString("mysql_port"),
		viper.GetString("mysql_database"),
	))
	if err != nil || dbCon.Ping() != nil {
		panic("failed to connect to database")
	}
	dbCon.SetConnMaxLifetime(time.Second)
	db = &DB{DB: dbCon}
}

func CloseDatabase() {
	if db != nil {
		err := db.DB.Close()
		if err != nil {
			fmt.Println("failed to close database connection")
		}
	}
}

func GetDB() *DB {
	return db
}

func (db DB) GetModelCount(table string, column string, value string) (int, error) {
	count := 0
	err := db.DB.Get(&count, "SELECT count(*) FROM "+table+" WHERE "+column+" = ?", value)
	return count, err
}
