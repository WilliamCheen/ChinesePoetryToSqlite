package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./poems.db")

	db, err := sql.Open("sqlite3", "./poems.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	content, err := ioutil.ReadFile("./create_table.sql")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := string(content)
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	var files []string

	sqlFileFolder := "./data/sql"

	err = filepath.Walk(sqlFileFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	for _, sqliFile := range files {
		content, err = ioutil.ReadFile(sqliFile)
		var sqlQuery = string(content)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(sqlQuery)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}
}
