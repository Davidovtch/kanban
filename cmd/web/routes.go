package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /login", app.getLoginPage)
	mux.HandleFunc("GET /register", app.getRegisterPage)
	mux.HandleFunc("GET /", app.getHomePage)
	mux.HandleFunc("GET /tasks", app.getTaskPage)
	mux.HandleFunc("GET /employee", app.getEmployeePage)
	// mux.HandleFunc("GET /taem", app.getTaemPage)
	mux.HandleFunc("GET /task/{id}", app.getUpdateTaskPage)
	mux.HandleFunc("GET /employee/{id}", app.getUpdateEmployeePage)
	// mux.HandleFunc("GET /logout", app.getLogout)

	mux.HandleFunc("POST /tasks", app.postTask)
	mux.HandleFunc("POST /employee", app.postEmployee)

	mux.HandleFunc("GET /delTask/{id}", app.delTask)
	mux.HandleFunc("GET /delEmployee/{id}", app.delEmployee)

	mux.HandleFunc("POST /upTask/{id}", app.updateTask)
	mux.HandleFunc("POST /upEmployee/{id}", app.updateEmployee)

	return mux
}
