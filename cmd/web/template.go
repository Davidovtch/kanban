package main

import (
	"log"
	"net/http"
	"text/template"
)

type pageData map[string]any

func render(w http.ResponseWriter, page string, data pageData) {
	t, err := template.ParseFiles("./assets/templates/" + page)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if data == nil {
		data = pageData{}
	}

	t.Execute(w, data)
}
