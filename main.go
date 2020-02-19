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

	port := viper.GetString("server_port")

	if len(port) == 0 {
		port = "80"
	}

	log.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	}

	_ = viper.BindEnv("mysql_username", "MYSQL_USERNAME")
	_ = viper.BindEnv("mysql_host", "MYSQL_HOST")
	_ = viper.BindEnv("mysql_password", "MYSQL_PASSWORD")
	_ = viper.BindEnv("mysql_port", "MYSQL_PORT")
	_ = viper.BindEnv("mysql_database", "MYSQL_DATABASE")
}

func main() {
	loadConfig()

	db.SetupDatabase()
	defer db.CloseDatabase()

	handleRequests()
}
