package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]                //Remove slash leading path
		t := templates.Lookup(requestedFile + ".html") //So Urls dont have to end in html all the time
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"

	//Every file in base path with html will be parsed
	template.Must(result.ParseGlob(basePath + "/*.html"))

	return result
}
