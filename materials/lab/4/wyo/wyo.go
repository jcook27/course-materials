package wyo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
}

var Assignments []Assignment
const Valkey string = "FooKey"

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response Response
	response.Assignments = Assignments
	params := mux.Vars(r)

	for i, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			assignment.Id =  r.FormValue("id")
			assignment.Title =  r.FormValue("title")
			assignment.Description =  r.FormValue("desc")
			assignment.Points, _ =  strconv.Atoi(r.FormValue("points"))
			Assignments[i] = assignment
			w.WriteHeader(http.StatusCreated)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}

type ResponseClasses struct{
	WyoClasses []WyoClass `json:"wyoclasses"`
}

type WyoClass struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
}

var WyoClasses []WyoClass
//const Valkey string = "FooKey"

func InitClass(){
	var WyoClass WyoClass
	WyoClass.Id = "COSC4010"
	WyoClass.Title = "Black Hat Go"
	WyoClass.Description = "Learn Go, the programming language not the board game, through cybersecurtiy related exercises."
	WyoClasses = append(WyoClasses, WyoClass)
}

//func APISTATUS(w http.ResponseWriter, r *http.Request) {
//	log.Printf("Entering %s end point", r.URL.Path)
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprintf(w, "API is up and running")
//}


func GetClassess(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response ResponseClasses

	response.WyoClasses = WyoClasses

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, WyoClass := range WyoClasses {
		if WyoClass.Id == params["id"]{
			json.NewEncoder(w).Encode(WyoClass)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, wyoclass := range WyoClasses {
			if wyoclass.Id == params["id"]{
				WyoClasses = append(WyoClasses[:index], WyoClasses[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response ResponseClasses
	response.WyoClasses = WyoClasses
	params := mux.Vars(r)

	for i, wyoclass := range WyoClasses {
		if wyoclass.Id == params["id"]{
			json.NewEncoder(w).Encode(wyoclass)
			wyoclass.Id =  r.FormValue("id")
			wyoclass.Title =  r.FormValue("title")
			wyoclass.Description =  r.FormValue("desc")
			WyoClasses[i] = wyoclass
			w.WriteHeader(http.StatusCreated)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var wyoclass WyoClass
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		wyoclass.Id =  r.FormValue("id")
		wyoclass.Title =  r.FormValue("title")
		wyoclass.Description =  r.FormValue("desc")
		WyoClasses = append(WyoClasses, wyoclass)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}