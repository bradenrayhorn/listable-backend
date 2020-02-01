package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

var lists []List

const RoutePrefix string = "/api"

func getLists(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(lists)
}

func handleRequests() {
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc(RoutePrefix + "/lists", getLists).Methods("GET")

  log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
  lists = []List{
    List{Title: "Walmart"},
  }

  handleRequests()
}
