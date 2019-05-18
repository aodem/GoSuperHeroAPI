package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type heroAttributes struct {
	PageID           int    `json:"pageid"`
	Name             string `json:"name"`
	Urlslug          string `json:"url"`
	ID               string `json:"id"`
	Align            string `json:"align"`
	Eye              string `json:"eye"`
	Hair             string `json:"hair"`
	Sex              string `json:"sex"`
	Gsm              string `json:"gsm"`
	Alive            string `json:"alive"`
	Appearances      int    `json:"appearances"`
	FirstAppearances string `json:"first appearance"`
	Year             int    `json:"year"`
}

type newCharacter struct {
	Name  string `json:"name"`
	Power string `json:"power"`
}

// AllCharacters : routing function for all characters
func AllCharacters(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:RVLmusic@tcp(127.0.0.1)/marvel")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM characters")
	if err != nil {
		log.Fatalln(err)
	}

	var charKeepr []heroAttributes

	for results.Next() {
		var theCharacters heroAttributes

		err := results.Scan(&theCharacters.PageID, &theCharacters.Name, &theCharacters.Urlslug, &theCharacters.ID, &theCharacters.Align, &theCharacters.Eye, &theCharacters.Hair, &theCharacters.Sex, &theCharacters.Gsm, &theCharacters.Alive, &theCharacters.Appearances, &theCharacters.FirstAppearances, &theCharacters.Year)
		if err != nil {
			log.Fatalln(err)
		}

		charKeepr = append(charKeepr, theCharacters)
	}

	byteArr, err := json.MarshalIndent(charKeepr, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "\n%v\n", string(byteArr))
}

// NewCharacter : start of the crud "read"
func NewCharacter(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:RVLmusic@tcp(127.0.0.1)/marvel")
	if err != nil {
		log.Fatalf("Connection error %v\n", err)
	}

	defer db.Close()

	if r.Method == "GET" {

		results, err := db.Query("SELECT name, power FROM new_characters")
		if err != nil {
			log.Fatalf("Query error: %v\n", err)
		}

		var newCollection []newCharacter

		for results.Next() {
			var character newCharacter

			err := results.Scan(&character.Name, &character.Power)
			if err != nil {
				log.Fatalln(err)
			}

			newCollection = append(newCollection, character)
		}

		byteArr, err := json.MarshalIndent(newCollection, "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(w, "\n%v\n", string(byteArr))

	} else if r.Method == "POST" {
		// "create" of crud
		decoder := json.NewDecoder(r.Body)
		var postC newCharacter
		err := decoder.Decode(&postC)
		if err != nil {
			panic(err)
		}

		var data = postC.Name
		var dataP = postC.Power

		add, err := db.Query("INSERT INTO new_characters (name, power) VALUES (?, ?)", data, dataP)
		if err != nil {
			log.Fatalln(err)
		}

		defer add.Close()

	}
}

// UpdateCharacter : "update" of crud
func UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["name"]
	fmt.Fprintf(w, "name: "+key)

	decoder := json.NewDecoder(r.Body)
	var postC newCharacter
	err := decoder.Decode(&postC)
	if err != nil {
		panic(err)
	}

	var data = postC.Name
	var dataP = postC.Power

	// database connection
	db, err := sql.Open("mysql", "root:RVLmusic@tcp(127.0.0.1)/marvel")
	if err != nil {
		log.Fatalln(err)
	}

	results, err := db.Query("UPDATE new_characters SET name=?, power=? WHERE name = ?", data, dataP, key)
	if err != nil {
		log.Println(err)
	}

	defer results.Close()
}

// DeleteCharacter : "delete" of crud
func DeleteCharacter(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)

	vars := mux.Vars(r)
	key := vars["name"]
	fmt.Fprintf(w, "name: "+key)

	fmt.Println(key)

	db, err := sql.Open("mysql", "root:RVLmusic@tcp(127.0.0.1)/marvel")
	if err != nil {
		log.Fatalln(err)
	}

	results, err := db.Query("DELETE FROM new_characters WHERE name = ?", key)
	if err != nil {
		log.Fatalln(err)
	}

	defer results.Close()

}
