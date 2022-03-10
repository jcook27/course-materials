package wyoclass

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	WyoClasses []WyoClass `json:"wyoclasses"`
}

type WyoClass struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
}

var WyoClasses []WyoClass
const Valkey string = "FooKey"

func InitClass(){
	var WyoClass WyoClass
	WyoClass.Id = "COSC4010-1"
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
	var response Response

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
	var response Response
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