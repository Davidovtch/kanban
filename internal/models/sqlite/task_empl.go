package sqlite

import (
	"database/sql"

	"github.com/davidovtch/Projeto-testes/internal/models"
)

type TEModel struct {
	DB *sql.DB
}

func (m *TEModel) All() ([]models.Task_Empl, error) {
	stmt := `SELECT id,task_id,employee_id FROM task_employee`

	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	relations := []models.Task_Empl{}
	for row.Next() {
		te := models.Task_Empl{}
		err := row.Scan(&te.ID, &te.Task_id, &te.Employee_id)
		if err != nil {
			return nil, err
		}

		relations = append(relations, te)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return relations, nil
}

func (m *TEModel) Find(id int) (models.Task_Empl, error) {
	var te models.Task_Empl

	stmt := `SELECT id,task_id,employee_id FROM task_employee WHERE id = ?`
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

func (m *TEModel) Delete_Task(id int) error {
	stmt := `DELETE FROM task_employee WHERE task_id = ?`

	_, err := m.DB.Exec(stmt, id)
	return err
}

func (m *TEModel) Delete_Employee(id int) error {
	stmt := `DELETE FROM task_employee WHERE employee_id = ?`

	_, err := m.DB.Exec(stmt, id)
	return err
}

func (m *TEModel) Update(task_id, empl_id, id int) error {
	stmt := `UPDATE task_employee SET task_id = $1,employee_id = $2 WHERE id = $3`

	_, err := m.DB.Exec(stmt, task_id, empl_id, id)
	return err
}

func (m *TEModel) Relations() ([]models.HomeView, error) {
	stmt := `SELECT te.id,t.id,t.name,t.status,t.endDate,e.id,e.name 
	FROM task_employee te INNER JOIN tasks t ON te.task_id = t.id 
	INNER JOIN employees e ON te.employee_id = e.id;`

	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	relations := []models.HomeView{}
	for row.Next() {
		te := models.HomeView{}
		err = row.Scan(&te.Relation_id, &te.Task_id, &te.Task_name, &te.Task_status, &te.Task_endDate, &te.Employee_id, &te.Employee_name)
		if err != nil {
			return nil, err
		}

		relations = append(relations, te)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return relations, nil
}
