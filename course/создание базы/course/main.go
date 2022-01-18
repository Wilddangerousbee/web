package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/lib/pq"
)

var connStr = "user=postgres password=Balonka1 dbname=postgres sslmode=disable port=5434"

func requestСoncatenation(parameters []string) string {
	request := " ("

	request = request + " " + parameters[0]

	for i := 1; i < len(parameters)/2; i++ {
		request = request + ", " + parameters[i]
	}

	request = request + ") " + "VALUES"

	request = request + " (" + parameters[len(parameters)/2]

	for i := len(parameters)/2 + 1; i < len(parameters); i++ {
		request = request + ", " + parameters[i]
	}

	request = request + ") "

	return request
}

func getData(nameTable string, nameColumn string) []string {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query("select " + nameColumn + " from " + nameTable)

	var rows = []string{}

	var p string

	for result.Next() {

		result.Scan(&p)

		rows = append(rows, p)
	}

	return rows
}

func postData(nameTable string, nameColumns ...string) {
	db, err := sql.Open("postgres", connStr)

	request := requestСoncatenation(nameColumns)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into " + nameTable + " " + request)

	fmt.Println(result)

	if err != nil {
		panic(err)
	}
}

func update(nameTable string, id string, nameColumns ...string) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var request = " "

	request += nameColumns[0] + " = " + nameColumns[len(nameColumns)/2]

	for i := 2; i < len(nameColumns)/2; i++ {
		request += ", " + nameColumns[i] + " = " + nameColumns[len(nameColumns)/2+i]
	}

	fmt.Printf(request)

	result, err := db.Exec("update " + nameTable + " set " + request + " where id = " + id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}

func delete(nameTable string, id string) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from " + nameTable + " where id = " + id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func main() {
	names := [10]string{"'Солнечная батарея SilaSolar 180Вт ( 5BB )'", "'Солнечная батарея SilaSolar 280Вт ( 5BB )'", "'Солнечная батарея SilaSolar 330Вт PERC ( 5BB )'", "'Солнечная батарея SilaSolar 370Вт PERC ( 5BB )'", "'Солнечная батарея Jinko Solar 560Вт (Bifacial)'", "'Солнечная батарея SilaSolar 30Вт ( 5BB )'", "'Солнечная батарея SilaSolar 50Вт ( 5BB )'", "'Солнечная батарея SilaSolar 100Вт ( 5BB )'", "'Солнечная батарея SilaSolar 200Вт ( 5BB )'", "'Солнечная батарея SilaSolar 400Вт PERC ( 5BB )'"}
	efficiency := [10]string{"0.1966", "0.1824", "0.2132", "0.218", "0.225", "0.225", "0.321", "0.423", "0.123", "0.253"}
	voltage := [10]string{"180", "240", "330", "370", "560", "30", "50", "100", "200", "400"}
	price := [10]string{"7062", "10890", "12342", "14322", "2772", "4950", "8184", "15114", "15246", "17688"}
	square := [10]string{"1.0064", "1.6335", "1.0164", "1.94432", "2.734074", "0.196", "0.37408", "0.67802", "0.13167", "1.98396"}

	for i := 0; i < 1000; i++ {
		postData("solar_panel", "name", "efficiency", "voltage", "price", "square", "quantity", names[rand.Intn(10)], efficiency[rand.Intn(10)], voltage[rand.Intn(10)], price[rand.Intn(10)], square[rand.Intn(10)], strconv.Itoa(rand.Intn(10)))
	}

	var f = getData("solar_panel", "name")
	// postData("solar_panel", "name", "efficiency", "voltage", "price", "square", "'ferst h'", "0.05", "342342", "423", "4324")
	// update("solar_panel", "2", "name", "voltage", "'todo'", "8912")
	// delete("solar_panel", "4")

	for _, v := range f {
		fmt.Printf(v, "\n")
	}
}
