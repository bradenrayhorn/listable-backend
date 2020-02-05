package main

import (
  "github.com/bradenrayhorn/listable-backend/db"
  "github.com/gorilla/mux"
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
  
  db.SetupDatabase()
  defer db.CloseDatabase()
  
  handleRequests()
}
