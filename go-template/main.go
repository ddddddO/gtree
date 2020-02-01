package main

import (
	"html/template"
	"log"
	"net/http"
)

type RenderData struct {
	Title       string
	Description string
}

func main() {
	var base = `
<html>
<head>
  <title>{{.Title}}</title>
</head>
<body>
  <h1>Decscription</h1>
  <p>{{.Description}}</p>
</body>
</html>
`

	tmpl := template.Must(template.New("tmpl").Parse(base))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		data := RenderData{
			Title:       "Hello World!",
			Description: "training go template",
		}

		if err := tmpl.Execute(w, data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	layoutTmpl := template.Must(template.ParseFiles("layouts/index.html"))
	http.HandleFunc("/a", func(w http.ResponseWriter, _ *http.Request) {
		err := layoutTmpl.Execute(w,
			struct {
				Text string
			}{
				Text: "file path pattern",
			},
		)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Println("server lanched")
	http.ListenAndServe(":8080", nil)
}
