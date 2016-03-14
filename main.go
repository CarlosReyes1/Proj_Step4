package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"os"
	"text/template"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {

	useroni := User{}

	templ, error := template.ParseFiles("tpl.gohtml") // Parse template file
	if error != nil {
		log.Fatalln(error)
	}

	error = templ.Execute(os.Stdout, nil)
	if error != nil {
		log.Fatalln(error)
	}

	cookie, err := req.Cookie("session-fino")

	useroni.FirstName = req.FormValue("FirstName")
	useroni.LastName = req.FormValue("LastName")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "_" + useroni.FirstName + "_" + useroni.LastName,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

	templ.ExecuteTemplate(res, "tpl.gohtml", useroni)

}
