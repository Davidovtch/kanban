package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/davidovtch/Projeto-testes/internal/forms"
	"github.com/davidovtch/Projeto-testes/internal/models"
)

type homeRender struct {
	task_id     int
	task_name   string
	task_status string
	task_date   string
	employees   []models.Employees
}

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

	task_empl, err := app.task_empl.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	home := []homeRender{}
	tmp := []models.Employees{}
	for _, ts_em := range task_empl {
		for _, task := range tasks {
			for _, employee := range empl {
				if ts_em.Task_id == task.ID && employee.ID == ts_em.Employee_id {
					tmp = append(tmp, employee)
				}
			}
			home = append(home, homeRender{task_id: task.ID, task_name: task.Name, employees: tmp})
		}
	}
	log.Println("TOTAL", home)
	render(w, "home.html", pageData{"Render": home})
}

func (app *app) getTaskPage(w http.ResponseWriter, r *http.Request) {
	render(w, "task.html", nil)
}

func (app *app) getEmployeePage(w http.ResponseWriter, r *http.Request) {
	render(w, "employee.html", nil)
}

func (app *app) getTaemPage(w http.ResponseWriter, r *http.Request) {
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

	render(w, "taem.html", pageData{"Tasks": tasks, "Employees": empl})
}

func (app *app) getUpdateTaskPage(w http.ResponseWriter, r *http.Request) {
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

func (app *app) getUpdateEmployeePage(w http.ResponseWriter, r *http.Request) {
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

	form := forms.New(r.PostForm)
	form.Required("name", "status", "endDate")
	form.MaxLenght("name", 50)

	if !form.Valid() {
		render(w, "task.html", pageData{"Form": form})
		return
	}

	err := app.task.Insert(
		r.PostForm.Get("name"),
		r.PostForm.Get("status"),
		r.PostForm.Get("endDate"),
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

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLenght("name", 30)
	form.MaxLenght("email", 60)
	form.Email("email")

	if !form.Valid() {
		render(w, "employee.html", pageData{"Form": form})
		return
	}

	err := app.empl.Insert(
		r.PostForm.Get("name"),
		r.PostForm.Get("email"),
		r.PostForm.Get("password"),
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *app) postTaem(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("task", "employee")

	if !form.Valid() {
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
		render(w, "taem.html", pageData{"Tasks": tasks, "Employees": empl, "Form": form})
		return
	}

	task_id, err := strconv.Atoi(r.PostForm.Get("task"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	employee_id, err := strconv.Atoi(r.PostForm.Get("employee"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = app.task_empl.Insert(
		task_id,
		employee_id,
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

	form := forms.New(r.PostForm)
	form.Required("name", "status", "endDate")
	form.MaxLenght("name", 50)

	if !form.Valid() {
		render(w, "upTask.html", pageData{"Form": form})
		return
	}

	if err = app.task.Update(r.PostForm.Get("name"), r.PostForm.Get("status"), r.PostForm.Get("endDate"), id); err != nil {
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

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLenght("name", 30)
	form.MaxLenght("email", 60)
	form.Email("email")

	if !form.Valid() {
		render(w, "upEmployee.html", pageData{"Form": form})
		return
	}

	if err = app.empl.Update(r.PostForm.Get("name"), r.PostForm.Get("email"), r.PostForm.Get("password"), id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
