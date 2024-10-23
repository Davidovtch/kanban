package sqlite

import (
	"database/sql"
	"log"

	"github.com/davidovtch/Projeto-testes/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeModel struct {
	DB *sql.DB
}

func (m *EmployeeModel) All() ([]models.Employees, error) {
	stmt := `SELECT id,name,email FROM employees ORDER BY ID`

	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	employees := []models.Employees{}
	for row.Next() {
		e := models.Employees{}
		err = row.Scan(&e.ID, &e.Name, &e.Email)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	err = row.Err()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (m *EmployeeModel) Find(id int) (models.Employees, error) {
	var e models.Employees

	stmt := `SELECT id,name,email,password FROM employees WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&e.ID, &e.Name, &e.Email, &e.Password)
	if err != nil {
		return models.Employees{}, err
	}

	return e, nil
}

func (m *EmployeeModel) Auth(email, password string) (int, error) {
	var (
		hashed []byte
		id     int
	)

	stmt := `SELECT id,password FROM employees WHERE email = ?`
	row := m.DB.QueryRow(stmt, email)

	err := row.Scan(&id, &hashed)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *EmployeeModel) Insert(name, email, password string) error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO employees (name,email,password) VALUES (?,?,?)`

	_, err = m.DB.Exec(stmt, name, email, passwd)
	return err
}

func (m *EmployeeModel) Delete(id int) error {
	stmt := `DELETE FROM employees WHERE id=?`
	_, err := m.DB.Exec(stmt, id)

	return err
}

func (m *EmployeeModel) Update(name, email, password string, id int) error {
	var hashed []byte
	var newPasswd []byte

	stmt := `UPDATE employees SET name = $1,email = $2,password = $3 WHERE id = $4`
	stmt2 := `SELECT password FROM employees WHERE id = ?`

	row := m.DB.QueryRow(stmt2, id)
	err := row.Scan(&hashed)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		log.Println("new password")
		newPasswd, err = bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			return err
		}
	} else {
		log.Println("same password as before")
		newPasswd = hashed
	}

	_, err = m.DB.Exec(stmt, name, email, newPasswd, id)

	return err
}
