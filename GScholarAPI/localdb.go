package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"time"
)

const datelayout string = "2006-01-02"

func openConnectiontoDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gscholarapi")
	if err!= nil {
		log.Fatal("Problem connecting to database", err)
		panic(err.Error())
	}

	return db
}

func insertRowintoDB(db *sql.DB, gscholarinfo GScholar) {
	querystmt, err := db.Prepare("INSERT gscholarinfo SET fullname=?, profilelink=?, citationscount=?, designation=?, university=?, created=?")
	if err!= nil {
		log.Fatal("Problem inserting data into database", err)
	}

	current_time := time.Now().Local()
	_, err = querystmt.Exec(gscholarinfo.FullName, gscholarinfo.GScholarProfileLink, gscholarinfo.CitationsCount, gscholarinfo.Designation, gscholarinfo.University, current_time.Format(datelayout))
	if err!= nil {
		log.Fatal("Problem inserting data into database", err)
	}

	fmt.Printf("GScholar Info for %s has been inserted into the database", gscholarinfo.FullName)
}

func readfromDB(db *sql.DB, fullname string) GScholar {
	var gscholarInfo GScholar
	var creationdate string
	err := db.QueryRow("SELECT fullname, profilelink, citationscount, designation, university, created FROM gscholarinfo WHERE fullname=?", fullname).Scan(&gscholarInfo.FullName, &gscholarInfo.GScholarProfileLink, &gscholarInfo.CitationsCount, &gscholarInfo.Designation, &gscholarInfo.University, &creationdate)
	if err!= nil {
		fmt.Println("No records were found in database, returning an empty object")
		return gscholarInfo
	}

	//Creating a time object today which holds todays date
	timenow := time.Now().Local()
	today, _ := time.Parse(datelayout, timenow.Format(datelayout))

	//Converting creationdate string object to a time object
	created, err := time.Parse(datelayout, creationdate)
	if err!= nil {
		log.Fatal("Error Parsing the creation date of the record from DB")
	}

	//Checking if the information is stale - older than 7 days
	if today.Sub(created).Hours()/24 < 7 {
		//it is less than 7 so sending the fetched data
		return gscholarInfo
	} else
	{
		//it is stale data so sending an emtpy object
		var emptygscholarInfo GScholar
		return emptygscholarInfo
	}
	}