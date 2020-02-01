package models

import "github.com/jinzhu/gorm"

type List struct {
  gorm.Model
  Title string `json:"title"`
}

func GetAllLists() []*List {
  lists := make([]*List, 0)
  db.Find(&lists)

  return lists
}
