package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/davidovtch/Projeto-testes/internal/models/sqlite"
)

type app struct {
	task      *sqlite.TaskModel
	empl      *sqlite.EmployeeModel
	task_empl *sqlite.TEModel
}

func main() {
	db, err := sql.Open("sqlite3", "file:app.db?_fk=true")
	if err != nil {
		log.Fatal("Error opening/creating the database\n", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name VARCHAR(255) NOT NULL,
    	status VARCHAR(255) NOT NULL,
    	endDate TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS employees(
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name VARCHAR(255) NOT NULL,
    	email VARCHAR(255) UNIQUE NOT NULL,
    	password VARCHAR(255) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS task_employee(
    	id INTEGER PRIMARY KEY AUTOINCREMENT,    
    	task_id INTEGER NOT NULL,
    	employee_id INTEGER NOT NULL,
    	FOREIGN KEY(task_id) REFERENCES tasks(id),
    	FOREIGN KEY(employee_id) REFERENCES employees(id)
		);`)
	if err != nil {
		log.Fatal("Error creating the tables\n", err)
	}

	app := app{
		task: &sqlite.TaskModel{
			DB: db,
		},
		empl: &sqlite.EmployeeModel{
			DB: db,
		},
		task_empl: &sqlite.TEModel{
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
