package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wyo/wyo"
)


func main() {
	wyo.InitAssignments()
	wyo.InitClass()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyo.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyo.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyo.GetAssignment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyo.DeleteAssignment).Methods("DELETE")		
	router.HandleFunc("/assignment", wyo.CreateAssignment).Methods("POST")	
	router.HandleFunc("/assignments/{id}", wyo.UpdateAssignment).Methods("PUT")

	router.HandleFunc("/classes", wyo.GetClassess).Methods("GET")
	router.HandleFunc("/class/{id}", wyo.GetClass).Methods("GET")
	router.HandleFunc("/class/{id}", wyo.DeleteClass).Methods("DELETE")		
	router.HandleFunc("/class", wyo.CreateClass).Methods("POST")	
	router.HandleFunc("/classes/{id}", wyo.UpdateClass).Methods("PUT")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}