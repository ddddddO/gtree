package main

import (
	"html/template"
	"log"
	"net/http"
)

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

type RenderData struct {
	Title       string
	Description string
}

func main() {
	tmpl := template.Must(template.New("tmpl").Parse(base))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := RenderData{
			Title:       "Hello World!",
			Description: "training go template",
		}

		if err := tmpl.Execute(w, data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Println("server lanched")
	http.ListenAndServe(":8080", nil)
}
