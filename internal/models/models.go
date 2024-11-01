package models

type Tasks struct {
	ID      int
	Name    string
	Status  string
	EndDate string
}

type Employees struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task_Empl struct {
	ID          int
	Task_id     int
	Employee_id int
}

type HomeView struct {
	Relation_id   int
	Task_id       int
	Task_name     string
	Task_status   string
	Task_endDate  string
	Employee_id   int
	Employee_name string
}
