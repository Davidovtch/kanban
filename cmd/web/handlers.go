package main

import (
	"net/http"
	"strconv"
)

func (app *app) getLoginPage(w http.ResponseWriter, r *http.Request) {
	render(w, "login.html", nil)
}

func (app *app) getRegisterPage(w http.ResponseWriter, r *http.Request) {
	render(w, "register.html", nil)
}

func (app *app) getHomePage(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.task.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	empl, err := app.empl.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	render(w, "home.html", pageData{"Tasks": tasks, "Employees": empl})
}

func (app *app) getTaskPage(w http.ResponseWriter, r *http.Request) {
	render(w, "task.html", nil)
}

func (app *app) getEmployeePage(w http.ResponseWriter, r *http.Request) {
	render(w, "employee.html", nil)
}

func (app *app) upTask(w http.ResponseWriter, r *http.Request) {
	tmp := r.PathValue("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	task, err := app.task.Find(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	render(w, "upTask.html", pageData{"Values": task})
}

func (app *app) upEmployee(w http.ResponseWriter, r *http.Request) {
	tmp := r.PathValue("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	empl, err := app.empl.Find(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	render(w, "upEmployee.html", pageData{"Values": empl})
}

/*
	POST
*/

func (app *app) postTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	name := r.PostForm.Get("name")
	status := r.PostForm.Get("status")
	endDate := r.PostForm.Get("endDate")

	err := app.task.Insert(
		name,
		status,
		endDate,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *app) postEmployee(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	passwd := r.PostForm.Get("password")

	err := app.empl.Insert(
		name,
		email,
		passwd,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

/*
	DELETE
*/

func (app *app) delTask(w http.ResponseWriter, r *http.Request) {
	tmp := r.PathValue("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = app.task.Delete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *app) delEmployee(w http.ResponseWriter, r *http.Request) {
	tmp := r.PathValue("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = app.empl.Delete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

/*
	PUT
*/

func (app *app) updateTask(w http.ResponseWriter, r *http.Request) {
	tmp := r.PathValue("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	name := r.PostForm.Get("name")
	status := r.PostForm.Get("status")
	date := r.PostForm.Get("endDate")

	if err = app.task.Update(name, status, date, id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *app) updateEmployee(w http.ResponseWriter, r *http.Request) {
	tmp := r.PathValue("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	passwd := r.PostForm.Get("password")

	if err = app.empl.Update(name, email, passwd, id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
