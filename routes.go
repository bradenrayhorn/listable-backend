package main

import (
	"github.com/bradenrayhorn/listable-backend/controllers"
	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {

	apiRouter := router.PathPrefix("/api").Subrouter()
	authRouter := router.PathPrefix("/api/auth").Subrouter()

	apiRouter.Use(controllers.AuthMiddleware)

	// Authentication
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")

	// Lists
	apiRouter.HandleFunc("/lists", controllers.GetAllLists).Methods("GET")

	// List Items
	apiRouter.HandleFunc("/list-items", controllers.AddListItem).Methods("POST")
	apiRouter.HandleFunc("/list-items", controllers.SetListItemChecked).Methods("PUT")

	// Groups
	apiRouter.HandleFunc("/groups", controllers.GetAllGroups).Methods("GET")
	apiRouter.HandleFunc("/groups", controllers.CreateGroup).Methods("POST")
}
