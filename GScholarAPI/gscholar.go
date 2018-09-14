package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type GScholar struct {
	FullName string                `json: "fullname"`
	GScholarProfileLink string     `json: "gscholarprofielink"`
	CitationsCount string          `json: "citationscount"`
	Designation string             `json: "designation"`
	University string              `json: "university"`
}

func extractProfileLinkDetails(goQuerydocument *goquery.Document) string {
	var gscholarProfileLink string
	goQuerydocument.Find(".gsc_oai_name a").Each(func(i int, s *goquery.Selection) {
		var Links []string
		Links = append(Links, "https://scholar.google.co.in")
		Link, _ := s.Attr("href")
		Links = append(Links, Link)
		gscholarProfileLink = strings.Join(Links, "")
	})

	return gscholarProfileLink

}

func extractFullName(goQuerydocument *goquery.Document) string{
	var fullName []string
	goQuerydocument.Find(".gsc_oai_name a").Each(func(i int, s *goquery.Selection) {
		fullName = append(fullName, s.Text())
	})
	return strings.Join(fullName, " ")
}

func extractDesignationUniversityDetails(goQuerydocument *goquery.Document) (string, string) {
	var designation, university string
	goQuerydocument.Find(".gsc_oai_aff").Each(func(i int, s *goquery.Selection) {
		description := string(s.Text())
		if strings.Contains(description, ", ") {
			p := strings.Split(s.Text(), ", ")
			designation, university = p[0], p[1]
		} else {
			designation = s.Text()
			university = "Info Not Available"
		}
	})
	return designation, university
}

func extractCitationsCount(goQuerydocument *goquery.Document) (string) {
	var citationsCount string
	goQuerydocument.Find(".gsc_oai_cby").Each(func(i int, s *goquery.Selection) {
		Citations := "0"
		if strings.Contains(s.Text(), "Cited by ") {
			runes := []rune(s.Text())
			Citations = string(runes[strings.Index(s.Text(), "Cited by ")+9:])
		}
		citationsCount = Citations
	})
	return citationsCount
}

func getContentByURL (htmlPageURL string) *goquery.Document {
	fmt.Println("Retreiving information from the following URL")
	fmt.Println(htmlPageURL)
	webPageContent, err :=http.Get(htmlPageURL)
	if err != nil {
		fmt.Println("Error retrieving data from Page ", htmlPageURL)
		log.Fatal(err)
	}

	defer webPageContent.Body.Close()

	goQuerydocument, err := goquery.NewDocumentFromReader(webPageContent.Body)
	if err != nil {
		log.Fatal("HTTP Response body load error: ", err)
	}

	// Debugging errors in Heroku
	/*contents, err := ioutil.ReadAll(webPageContent.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println(string(contents)[:200])*/
	return goQuerydocument
}

func getCompleteURLbyName(fullName string) string {

	//preparing full name to be sent as a query string
	parsedfullName := strings.Replace(fullName, " ", "+", -1)

	//https://scholar.google.com/citations?hl=en&view_op=search_authors&mauthors=Prasanta+Achanta&btnG=
	urlPattern := "https://scholar.google.com/citations?hl=en&view_op=search_authors&mauthors=%s&btnG="

	//forming URL string using fullname
	urlSearchByName := fmt.Sprintf(urlPattern,parsedfullName)

	return urlSearchByName
}

func getGoogleScholarDetails(fullName string) GScholar{

	urlToSearch := getCompleteURLbyName(fullName)
	goQueryDocument := getContentByURL(urlToSearch)

	name := extractFullName(goQueryDocument)
	profileLink :=	extractProfileLinkDetails(goQueryDocument)
	citationsCount := extractCitationsCount(goQueryDocument)
	designation, university := extractDesignationUniversityDetails(goQueryDocument)

	GScholarDetails := GScholar{name, profileLink, citationsCount, designation, university}

	return GScholarDetails
}