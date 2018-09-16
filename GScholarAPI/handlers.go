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

	//First search in local database
	db := openConnectiontoDB()
	defer db.Close()

	scholarDetails := readfromDB(db, fullName)

	//If the information is stale or not available try scraping the google scholar web page
	if scholarDetails.FullName == "" {
		fmt.Printf("Searching for name %s on google scholar\n", fullName)
		scholarDetails = getGoogleScholarDetails(fullName)
		//Write the fetched details to local DB
		insertRowintoDB(db, scholarDetails)
	}

	fmt.Println("Found these details: ")
	fmt.Println(scholarDetails)

	//Call template execute function to apply parsed template to data followed by write to output writer
	homePgt.Execute(w, scholarDetails)

	//json.NewEncoder(w).Encode("Hello! This page is under construction")
}

func APIHandler(w http.ResponseWriter, r * http.Request) {
	requestVars := mux.Vars(r)
	fullName := requestVars["name"]

	//First search in local database
	db := openConnectiontoDB()
	defer db.Close()

	scholarDetails := readfromDB(db, fullName)

	//If the information is stale or not available try scraping the google scholar web page
	if scholarDetails.FullName == "" {
		fmt.Printf("Searching for name %s on google scholar\n", fullName)
		scholarDetails = getGoogleScholarDetails(fullName)
		//Write the fetched details to local DB
		insertRowintoDB(db, scholarDetails)
	}

	json.NewEncoder(w).Encode(scholarDetails)
}