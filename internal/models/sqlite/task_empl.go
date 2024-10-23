package sqlite

import (
	"database/sql"

	"github.com/davidovtch/Projeto-testes/internal/models"
)

type TEModel struct {
	DB *sql.DB
}

// func (m *TEModel) All() (models.Tasks, models.Employees, error) {
// 	stmt := `SELECT `
// send the task_empl model and use the html to choose how to render the tasks
// and employees

// }

func (m *TEModel) Find(id int) (models.Task_Empl, error) {
	var te models.Task_Empl

	stmt := `SELECT task_id,employee_id FROM task_employee WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&te.ID, &te.Task_id, &te.Employee_id)
	if err != nil {
		return models.Task_Empl{}, err
	}

	return te, nil
}

func (m *TEModel) Insert(task_id, empl_id int) error {
	stmt := `INSERT INTO task_employee(task_id,employee_id) VALUES(?,?)`

	_, err := m.DB.Exec(stmt, task_id, empl_id)
	return err
}

func (m *TEModel) Delete(id int) error {
	stmt := `DELETE FROM task_employee WHERE id = ?`

	_, err := m.DB.Exec(stmt, id)
	return err
}

func (m *TEModel) Update(task_id, empl_id, id int) error {
	stmt := `UPDATE FROM task_employee SET task_id = $1,employee_id = $2 WHERE id = $3`

	_, err := m.DB.Exec(stmt, task_id, empl_id, id)
	return err
}
