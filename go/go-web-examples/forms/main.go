package main

// ref https://gowebexamples.com/forms/

import (
	"html/template"
	"log"
	"net/http"
)

type Account struct {
	Email    string
	UserName string
	Password string
}

func main() {
	http.HandleFunc("/", render)

	log.Println("server launched")
	http.ListenAndServe(":8080", nil)
}

var tmpl = template.Must(template.ParseFiles("form.html"))

func render(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := tmpl.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	account := Account{
		Email:    r.FormValue("email"),
		UserName: r.FormValue("userName"),
		Password: r.FormValue("password"),
	}

	err := tmpl.Execute(w, struct {
		Logined  bool
		UserName string
	}{
		true,
		account.UserName,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
