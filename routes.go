package main

import (
  "github.com/bradenrayhorn/listable-backend/controllers"
  "github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {

  router.HandleFunc("/api/lists", controllers.GetAllLists).Methods("GET")

}
