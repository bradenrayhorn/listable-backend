package main

import (
  "github.com/bradenrayhorn/listable-backend/models"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/spf13/viper"
  "log"
  "net/http"
)

func handleRequests() {
  router := mux.NewRouter().StrictSlash(true)

  registerRoutes(router)

  log.Fatal(http.ListenAndServe(":"+viper.GetString("server_port"), router))
}

func loadConfig() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")

  if err := viper.ReadInConfig(); err != nil {
    panic("failed to read config")
  }
}

func main() {
  loadConfig()

  models.SetupDatabase()
  defer models.CloseDatabase()

  handleRequests()
}
