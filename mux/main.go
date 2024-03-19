package main

import (
	"fmt"
	"log"
	"mux/controller"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mhsMux", controller.InsertMhsMux).Methods("POST")
	router.HandleFunc("/mhsMux", controller.GetMhsMux).Methods("GET")
	router.HandleFunc("/mhsMux/{id}", controller.UpdateMhsMux).Methods("PUT")
	router.HandleFunc("/mhsMux/{id}", controller.DeleteMhsMux).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
