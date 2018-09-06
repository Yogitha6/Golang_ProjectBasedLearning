package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Collaborator struct {
	FullName string
	Location string
}

var Collaborators []Collaborator

func sanitizeData(inputtagdata string) string {
	inputtagdataLowerCase := strings.ToLower(inputtagdata)
	switch {
	case strings.Contains(inputtagdataLowerCase, "student"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "research"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "post-doc"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "faculty"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "scholar"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "copec"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "co-pec"): fallthrough
	case strings.Contains(inputtagdataLowerCase, "graduate"):
		return ""
	}
	//removing Dr. & Prof. Titles
	if strings.Contains(inputtagdata, "Dr. ") {
		runes := []rune(inputtagdata)
		inputtagdata = string(runes[strings.Index(inputtagdata, "Dr. ")+4:])
	}
	if strings.Contains(inputtagdata, "Prof. ") {
		runes := []rune(inputtagdata)
		inputtagdata = string(runes[strings.Index(inputtagdata, "Prof. ")+6:])
	}
	if strings.Contains(inputtagdata, "Mr. ") {
		runes := []rune(inputtagdata)
		inputtagdata = string(runes[strings.Index(inputtagdata, "Mr. ")+4:])
	}
	if strings.Contains(inputtagdata, "Mrs. ") {
		runes := []rune(inputtagdata)
		inputtagdata = string(runes[strings.Index(inputtagdata, "Mr. ")+4:])
	}
	return inputtagdata
}

func findLocation(personName string) (location string) {
	/* Need to replace the code with an API call which will provide the required location information */
	location = "Denver"
	return
}

/* */
func processSelection(index int, tagdata *goquery.Selection) {
	if fullname := sanitizeData(tagdata.Text()); fullname!="" {
		fmt.Println(fullname)
		Collaborators = append(Collaborators, Collaborator{fullname, findLocation(fullname)})
	}
}

func main() {
	/* net/http library provides a Get call to send a GET request to any URL. It returns a pointer to the object of Response struct and an error */
	webPageContent, err :=http.Get("http://ecee.colorado.edu/~maksimov/students.html")
	if err != nil {
		log.Fatal(err)
	}

	/*The http Response object has a member body which is streaming data on-demand from the server. The HTTP/1.X "keep-alive" TCP connections might not
	reusable if entire body content is not read followed by a close of stream
	 */
	 defer webPageContent.Body.Close()

	 /* goquery package provides us with DOM parsing & manipulation abilities similar to jQuery, we will use that library
	 inorder to extract the text in <strong> & <b> html tags *?
	  */
	  goQuerydocument, err := goquery.NewDocumentFromReader(webPageContent.Body)
	  if err != nil {
	  	log.Fatal("HTTP Response body load error: ", err)
	  }

	  goQuerydocument.Find("strong").Each(processSelection)
}