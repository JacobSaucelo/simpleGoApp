package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		app := template.Must(template.ParseFiles("./app/index.html"))
		app.Execute(res, nil)
	})

	http.HandleFunc("/posts/", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			post := getPosts()
			tmpl := template.Must((template.ParseFiles("./app/index.html")))
			tmpl.ExecuteTemplate(res, "posts-element", post)
		}
	})

	log.Println("Running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
