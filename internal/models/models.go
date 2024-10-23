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
