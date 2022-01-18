package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var nameCur = ""
var passwordCur = ""
var name = &nameCur
var password = &passwordCur

func home(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("ui/html/home.html"))
	var tpl1 = template.Must(template.ParseFiles("ui/html/inputSwihtMail.html"))
	var tpl2 = template.Must(template.ParseFiles("ui/html/inputS.html"))

	flUsername := false
	flMail := false

	nameH := r.FormValue("mail")
	passwordH := r.FormValue("password")

	fmt.Printf("Name home: " + nameH)

	nameDB := getData("users", "username")
	idDB := getData("users", "id")
	passwordDB := getData("users", "password")
	userInfoId := getData("user_info", "id")

	for i := 0; i < len(nameDB); i++ {
		if nameH == nameDB[i] && passwordH == passwordDB[i] {
			flUsername = true
			for _, v := range userInfoId {
				if v == idDB[i] {
					flMail = true
				}
			}
		}
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if nameH == "" || passwordH == "" {
		tpl.Execute(w, nil)
	} else if flUsername == true {
		nameCur = nameH
		if flMail == true {
			tpl2.Execute(w, nil)
			return
		}
		tpl1.Execute(w, nil)
	} else if flUsername == false {
		w.Write([]byte("пользователя c такой парой логин пароль не существует"))
	}

	fmt.Println(name, password)
}

func login(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("ui/html/login.html"))
	var tpl1 = template.Must(template.ParseFiles("ui/html/home.html"))

	flUsername := true

	var nameH = r.FormValue("mail")
	var passwordH = r.FormValue("password")

	nameDB := getData("users", "username")

	for i := 0; i < len(nameDB); i++ {
		if nameH == nameDB[i] {
			flUsername = false
		}
	}

	if nameH == "" || passwordH == "" {
		tpl.Execute(w, nil)
		return
	} else if flUsername == false {
		w.Write([]byte("пользователь с таким именем уже есть"))
		return
	} else {
		postData("users", "username", "password", "'"+nameH+"'", "'"+passwordH+"'")
		tpl1.Execute(w, nil)
	}
}

func calculateS(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("ui/html/inputSwihtMail.html"))
	var tpl1 = template.Must(template.ParseFiles("ui/html/inputS.html"))
	var sSendMail string
	var userId string

	usersN := getData("users", "username")
	usersId := getData("users", "id")
	userInfoId := getData("user_info", "id")
	userInfoMail := getData("user_info", "mail")
	haveMail := ""

	fmt.Printf("NAME calculate" + *name)

	for i := 0; i < len(usersN); i++ {
		if usersN[i] == *name {
			userId = usersId[i]
		}
	}

	for i := 0; i < len(userInfoId); i++ {
		if userInfoId[i] == userId {
			haveMail = userInfoMail[i]
		}
	}

	mail := r.FormValue("mail")
	square, err := strconv.ParseFloat(r.FormValue("square"), 64)

	if mail != "" {
		if haveMail == "" {
			postData("user_info", "id", "mail", userId, "'"+mail+"'")
		}
	}
	if haveMail != "" {
		mail = haveMail
	}

	if err != nil {
		if haveMail == "" {
			tpl.Execute(w, nil)
		} else {
			tpl1.Execute(w, nil)
		}

	} else {
		var res = updatePerfectCombination(square)

		fmt.Printf("10")
		for i := 0; i < len(res); i++ {
			sSendMail += "<p>" + res[i].Name + " " + string(res[i].Quantity) + " " + fmt.Sprintf("%f", res[i].Square) + " " + fmt.Sprintf("%f", res[i].UtilityCoefficient) + "<p>"
		}
		body := `
		<html>
		<body>
		<h3>
		` + sSendMail + `
		</h3>
		</body>
		</html>
		`
		subject := "Используйте Golang для отправки почты"

		if mail != "" {
			SendToMail(mail, subject, body, "html")
		}

		w.Write([]byte(body))
	}

}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/calculateS", calculateS)

	http.ListenAndServe(":"+port, mux)
}
