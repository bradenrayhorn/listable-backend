package models

type List struct {
  AutoId
  Title string `json:"title"`
  ModelTimestamps
}

func GetAllLists() []*List {
  lists := make([]*List, 0)
  db.Find(&lists)

  return lists
}
