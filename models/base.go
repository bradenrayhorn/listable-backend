package models

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "github.com/spf13/viper"
  "time"
)

var db *gorm.DB

func SetupDatabase() {
  dbCon, err := gorm.Open("mysql",
    fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
      viper.GetString("mysql_username"),
      viper.GetString("mysql_password"),
      viper.GetString("mysql_host"),
      viper.GetString("mysql_port"),
      viper.GetString("mysql_database"),
    ))
  if err != nil {
    panic("failed to connect to database.")
    return
  }
  db = dbCon

  db.AutoMigrate(&List{})
}

func CloseDatabase() {
  if db != nil {
    db.Close()
  }
}

type AutoId struct {
  ID uint `gorm:"primary_key" json:"id"`
}

type ModelTimestamps struct {
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
}
