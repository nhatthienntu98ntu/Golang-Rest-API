package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/config"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/router"
)

const (
	DEFAULT_HOST = "localhost"
	port         = "8000"
)

func main() {
	// Create new router and config router
	r := mux.NewRouter()
	router.RegisterUsersRouter(r)
	srv := &http.Server{
		Handler: r,
		Addr:    DEFAULT_HOST + ":" + port,
	}

	// Connect and Close connect DB
	config.Connect()
	defer config.CloseDB()
	// Run port 8000
	fmt.Println("Start port: 8000")
	log.Fatal(srv.ListenAndServe())
}
