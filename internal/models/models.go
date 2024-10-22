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
