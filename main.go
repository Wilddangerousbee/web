package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type text struct {
	Text string
}

func main() {
	var password string
	var login string

	http.HandleFunc("/angular", func(w http.ResponseWriter, r *http.Request) {
		textRespons := []text{}

		for i := 0; i < 10; i++ {
			textRespons = append(textRespons, text{"Привет" + strconv.Itoa(i)})
		}

		t, _ := json.Marshal(textRespons)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		fmt.Println(r.URL.Path)
		fmt.Println(r.Body)
		w.Write(t)
	})

	http.HandleFunc("/angular/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		body, _ := ioutil.ReadAll(r.Body)

		fmt.Println("Response\n", string(body))

		login, password = strings.Split(string(body), " ")[0], strings.Split(string(body), " ")[1]

		fmt.Println(r.URL.Path)

		var req = text{Text: "OK"}

		if password == "Balonka1" && login == "akozachenko" {
			reqJSON, _ := json.Marshal(req)
			w.Write(reqJSON)
		} else {
			req.Text = "NO"
			reqJSON, _ := json.Marshal(req)
			w.Write([]byte(reqJSON))
		}

		fmt.Println(login)
		fmt.Println(password)
	})

	http.HandleFunc("/angular/login/true", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		fmt.Println(r.URL.Path)

		var req = text{Text: "OK"}

		if password == "Balonka1" && login == "akozachenko" {
			reqJSON, _ := json.Marshal(req)
			w.Write(reqJSON)
		} else {
			req.Text = "NO"
			reqJSON, _ := json.Marshal(req)
			w.Write([]byte(reqJSON))
		}

		fmt.Println(login)
		fmt.Println(password)
	})

	http.ListenAndServe(":8080", nil)
}
