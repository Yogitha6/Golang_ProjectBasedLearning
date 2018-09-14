package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

func HomePageGETHandler(w http.ResponseWriter, r *http.Request) {

	homePgt := template.New("homepage")
	homePgt, _ = homePgt.Parse(homePagetmpl)

	fmt.Println(r.Method)

	//Call template execute function to apply parsed template to data followed by write to output writer
	fmt.Println("Method not Post, returning empty data")
	homePgt.Execute(w, "")
}

func HomePagePOSTHandler(w http.ResponseWriter, r *http.Request) {
	homePgt := template.New("homepage")
	//After creating a new template object, use it to parse a string defined above
	homePgt, _ = homePgt.Parse(homePagetmpl)

	fmt.Println(r.Method)

	fullName := r.FormValue("fullName")
	fmt.Printf("Searching for name %s on google scholar\n", fullName)

	scholarDetails := getGoogleScholarDetails(fullName)
	fmt.Println("Found these details: ")
	fmt.Println(scholarDetails)

	//Call template execute function to apply parsed template to data followed by write to output writer
	homePgt.Execute(w, scholarDetails)

	//json.NewEncoder(w).Encode("Hello! This page is under construction")
}

func APIHandler(w http.ResponseWriter, r * http.Request) {
	requestVars := mux.Vars(r)
	scholarDetails := getGoogleScholarDetails(requestVars["name"])
	json.NewEncoder(w).Encode(scholarDetails)
}