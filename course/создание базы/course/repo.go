// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// var connStr = "user=postgres password=Balonka1 dbname=postgres sslmode=disable"

// func requestСoncatenation(parameters []string) string {
// 	request := " ("

// 	request = request + " " + parameters[0]

// 	for i := 1; i < len(parameters)/2; i++ {
// 		request = request + ", " + parameters[i]
// 	}

// 	request = request + ") " + "VALUES"

// 	request = request + " (" + parameters[len(parameters)/2]

// 	for i := len(parameters)/2 + 1; i < len(parameters); i++ {
// 		request = request + ", " + parameters[i]
// 	}

// 	request = request + ") "

// 	return request
// }

// func getData(nameTable string, nameColumn string) []string {
// 	db, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	result, err := db.Query("select " + nameColumn + " from " + nameTable)

// 	var rows = []string{}

// 	var p string

// 	for result.Next() {

// 		result.Scan(&p)

// 		rows = append(rows, p)
// 	}

// 	return rows
// }

// func postData(nameTable string, nameColumns ...string) {
// 	db, err := sql.Open("postgres", connStr)

// 	request := requestСoncatenation(nameColumns)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	result, err := db.Exec("insert into " + nameTable + " " + request)

// 	fmt.Println(result)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func update(nameTable string, id string, nameColumns ...string) {
// 	db, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	var request = " "

// 	request += nameColumns[0] + " = " + nameColumns[len(nameColumns)/2]

// 	for i := 2; i < len(nameColumns)/2; i++ {
// 		request += ", " + nameColumns[i] + " = " + nameColumns[len(nameColumns)/2+i]
// 	}

// 	fmt.Printf(request)

// 	result, err := db.Exec("update " + nameTable + " set " + request + " where id = " + id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(result.RowsAffected())
// }

// func delete(nameTable string, id string) {

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	result, err := db.Exec("delete from " + nameTable + " where id = " + id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(result.RowsAffected()) // количество удаленных строк
// }
