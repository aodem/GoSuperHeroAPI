package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type heroAttributes struct {
	PageID           int
	Name             string
	Urlslug          string
	ID               string
	Align            string
	Eye              string
	Hair             string
	Sex              string
	Gsm              string
	Alive            string
	Appearances      int
	FirstAppearances string
	Year             int
}

//Retrieve : acessing data
func Retrieve() {
	db, err := sql.Open("mysql", "root:RVLmusic@tcp(127.0.0.1)/marvel")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM characters")
	if err != nil {
		log.Fatalln(err)
	}

	defer results.Close()

	for results.Next() {

		var theCharacters heroAttributes

		err := results.Scan(&theCharacters.PageID, &theCharacters.Name, &theCharacters.Urlslug, &theCharacters.ID, &theCharacters.Align, &theCharacters.Eye, &theCharacters.Hair, &theCharacters.Sex, &theCharacters.Gsm, &theCharacters.Alive, &theCharacters.Appearances, &theCharacters.FirstAppearances, &theCharacters.Year)
		if err != nil {
			log.Fatalln(err)
		}

		byteArr, err := json.MarshalIndent(theCharacters, "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(res, "\n%v\n", string(byteArr))
	}
}
