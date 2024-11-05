package tests

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/davidovtch/Projeto-testes/internal/models/sqlite"
)

type test_app struct {
	task      *sqlite.TaskModel
	empl      *sqlite.EmployeeModel
	task_empl *sqlite.TEModel
}

var app *test_app

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", "file:test.db?_fk=true")
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	app = &test_app{
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

	m.Run()
}

func TestAddTask(t *testing.T) {
	err := app.task.Insert("TestAdd", "todo", "2024-11-05")
	if err != nil {
		t.Error("Expected to succeed adding a task,instead got error: ", err)
	}
}

func TestTaskStatusUpdate(t *testing.T) {
	name := "Update Status"
	firstStatus := "todo"
	date := "2024-11-05"
	newStatus := "done"

	err := app.task.Insert(name, firstStatus, date)
	if err != nil {
		t.Error("Failed to insert task to be updated", err)
	}

	err = app.task.Update(name, newStatus, date, 2)
	if err != nil {
		t.Error("Expected to succeed updating a task status,instead got error: ", err)
	}
}

func TestRelationshipCreation(t *testing.T) {
	if err := app.empl.Insert("rEmployee", "rel@mail.com", "123"); err != nil {
		t.Error("Failed to create employee ", err)
	}

	if err := app.task.Insert("rTask", "doing", "2024-11-05"); err != nil {
		t.Error("Failed to create task ", err)
	}

	if err := app.task_empl.Insert(3, 1); err != nil {
		t.Error("Expected to succeed creating a relationship,instead got error: ", err)
	}
}
