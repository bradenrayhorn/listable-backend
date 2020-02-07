package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"reflect"
	"strings"
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

func getFields(v reflect.Value, fields []string, values []interface{}, params []string) ([]string, []interface{}, []string) {
	modelType := v.Type()
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		if field.Anonymous {
			// Composite field
			fields, values, params = getFields(v.Field(i), fields, values, params)
		} else {
			// Normal field
			if !v.Field(i).IsZero() {
				switch v.Field(i).Type().String() {
				case "uint":
					values = append(values, v.Field(i).Uint())
				case "string":
					values = append(values, v.Field(i).String())
				}
				fields = append(fields, strcase.ToSnake(field.Name))
				params = append(params, "?")
			}
		}
	}
	return fields, values, params
}

func (db DB) Create(m interface{}) {
	table := m.(HasTableName).TableName()
	query := "insert into " + table + "("
	// Loop all model fields
	fields, values, params := getFields(reflect.ValueOf(m).Elem(), make([]string, 0), make([]interface{}, 0), make([]string, 0))
	query += strings.Join(fields, ",") + ") values (" + strings.Join(params, ",") + ")"
	r, err := db.DB.Exec(query, values...)
	if err != nil {
		fmt.Println(err) // TODO return error
	} else {
		if id, err := r.LastInsertId(); err == nil && reflect.ValueOf(m).Elem().FieldByName("Id").IsValid() {
			reflect.ValueOf(m).Elem().FieldByName("Id").SetUint(uint64(id))
		}
	}
}

func (db DB) GetModelCount(table string, column string, value string) (int, error) {
	count := 0
	err := db.DB.Get(&count, "SELECT count(*) FROM "+table+" WHERE "+column+" = ?", value)
	return count, err
}
