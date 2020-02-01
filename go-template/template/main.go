package main
//ref https://gowebexamples.com/templates/

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

	http.HandleFunc("/b", renderThirdHTML)

	log.Println("server lanched")
	http.ListenAndServe(":8080", nil)
}

type ThirdData struct {
	Title string
	Todos []Todo
}

type Todo struct {
	Title       string
	Description string
	IsDone      bool
}

var thirdTmpl = template.Must(template.ParseFiles("layouts/third.html"))

func renderThirdHTML(w http.ResponseWriter, _ *http.Request) {
	data := ThirdData{
		Title: "Third html redered!",
		Todos: []Todo{
			Todo{Title: "refacter working process", Description: "その名の通り", IsDone: false},
			Todo{Title: "棚卸", Description: "...", IsDone: false},
			Todo{Title: "hobby", Description: "hobby", IsDone: true},
		},
	}

	err := thirdTmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
