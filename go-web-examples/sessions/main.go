package main

// https://gowebexamples.com/sessions/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("0123456789abcdef")
	store = sessions.NewCookieStore(key)
)

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "input-cookie-value")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 認証済みとして登録
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "input-cookie-value")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 認証済みから消去
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func recieveSecretMsg(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "input-cookie-value")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 認証済みかチェック
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Forbidden...", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func main() {
	http.HandleFunc("/login", logging(login))
	http.HandleFunc("/logout", logging(logout))
	http.HandleFunc("/secret", logging(recieveSecretMsg))

	fmt.Println("server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// https://gowebexamples.com/basic-middleware/
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("access to %s", r.URL.Path)
		f(w, r)
	}
}
