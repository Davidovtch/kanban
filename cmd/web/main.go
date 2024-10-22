package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/davidovtch/Projeto-testes/internal/models/sqlite"
)

type app struct {
	task *sqlite.TaskModel
	empl *sqlite.EmployeeModel
}

func main() {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		task: &sqlite.TaskModel{
			DB: db,
		},
		empl: &sqlite.EmployeeModel{
			DB: db,
		},
	}

	srv := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listening on http://localhost:8000")
	log.Fatal(srv.ListenAndServe())
}
