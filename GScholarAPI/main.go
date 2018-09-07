package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var homePagetmpl = `<html>
<head>
<title>Google Scholar API </title>
</head>
<body>
<h1> Google Scholar API </h1>
<!-- <p> This page allows you to upload a file (containing full names of users separated by newline) to retrieve google scholar meta data </p> -->
<p> Type a full name in the below text box to retrieve related google scholar meta data </p>
<form method="POST">
<label>Full Name: </label><br/>
<input type="text" name="fullName"><br />
<input type="submit">
</form>

<h3> Details are: </h3>  
    Full Name: &emsp; {{ .FullName }} <br/>
    Google Scholar Profile Link : &emsp; <a href="{{ .GScholarProfileLink }}"> {{ .FullName }} </a> <br/>
	Citations Count: &emsp; {{ .CitationsCount }} <br/>
	Designation: &emsp; {{ .Designation }} <br/>
	University/Affliation: &emsp; {{ .University }} <br/>
</body>
</html>
`

type GScholar struct {
	FullName string
	GScholarProfileLink string
	CitationsCount string
	Designation string
	University string
	//websiteLink string
}

/*func getProfileURL(givenURL string ) string {

	return ""
}*/

func extractScholarDetails(htmlPageURL string) GScholar{

	var GScholarDetails GScholar
	webPageContent, err :=http.Get(htmlPageURL)
	if err != nil {
		log.Fatal(err)
	}

	defer webPageContent.Body.Close()

	goQuerydocument, err := goquery.NewDocumentFromReader(webPageContent.Body)
	if err != nil {
		log.Fatal("HTTP Response body load error: ", err)
	}

	goQuerydocument.Find(".gsc_oai_name a").Each(func(i int, s *goquery.Selection) {
		var Links []string
		Links = append(Links, "https://scholar.google.co.in")
		Link, _ := s.Attr("href")
		Links = append(Links, Link)
		GScholarDetails.GScholarProfileLink = strings.Join(Links, "")
		})

	var fullName []string
	goQuerydocument.Find(".gsc_oai_name a").Each(func(i int, s *goquery.Selection) {
		fullName = append(fullName, s.Text())
		})
	GScholarDetails.FullName = strings.Join(fullName, " ")

	goQuerydocument.Find(".gsc_oai_aff").Each(func(i int, s *goquery.Selection) {
		description := string(s.Text())
		if strings.Contains(description, ",") {
			p := strings.Split(s.Text(), ", ")
			Designation, University := p[0], p[1]
			GScholarDetails.Designation = Designation
			GScholarDetails.University = University
		} else {
			GScholarDetails.Designation = s.Text()
			GScholarDetails.University = "Info Not Available"
		}
		})

	goQuerydocument.Find(".gsc_oai_cby").Each(func(i int, s *goquery.Selection) {
		Citations := "0"
		if strings.Contains(s.Text(), "Cited by ") {
			runes := []rune(s.Text())
			Citations = string(runes[strings.Index(s.Text(), "Cited by ")+9:])
		}
		GScholarDetails.CitationsCount = Citations
		})

	return GScholarDetails
}

func getGoogleScholarDetails(fullName string) GScholar{

	//preparing full name to be sent as a query string
	parsedfullName := strings.Replace(fullName, " ", "+", -1)

	//https://scholar.google.com/citations?hl=en&view_op=search_authors&mauthors=Prasanta+Achanta&btnG=
	urlPattern := "https://scholar.google.com/citations?hl=en&view_op=search_authors&mauthors=%s&btnG="

	//forming URL string using fullname
	urlSearchByName := fmt.Sprintf(urlPattern,parsedfullName)

	//Make URL of the final profile page for advanced details please include this (out of scope as of now)
	//profileURL := getProfileURL(urlSearchByName)

	//Extract the required details using this profile page URL
	gScholarDetails := extractScholarDetails(urlSearchByName)

	return gScholarDetails
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	homePgt := template.New("homepage")
	//After creating a new template object, use it to parse a string defined above
	homePgt, _ = homePgt.Parse(homePagetmpl)

	if r.Method != http.MethodPost {
		//Call template execute function to apply parsed template to data followed by write to output writer
		homePgt.Execute(w, "")
		return
	}

	fullName := r.FormValue("fullName")

	scholarDetails := getGoogleScholarDetails(fullName)
	homePgt.Execute(w, scholarDetails)

	//json.NewEncoder(w).Encode("Hello! This page is under construction")
}

func APIHandler(w http.ResponseWriter, r * http.Request) {
	json.NewEncoder(w).Encode("This page is under construction")
}

func main() {
	// Gorilla mux package - HTTP Request multiplexer helps in request routing and dispatching to handlers
	router := mux.NewRouter()
	router.HandleFunc("/", HomePageHandler)
	router.HandleFunc("/gscholarapi/", APIHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}