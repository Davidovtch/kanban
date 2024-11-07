package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/davidovtch/Projeto-testes/internal/models"
)

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) All() ([]models.Tasks, error) {
	stmt := `SELECT id,name,status,endDate FROM tasks ORDER BY id`

	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	tasks := []models.Tasks{}
	for row.Next() {
		t := models.Tasks{}
		err := row.Scan(&t.ID, &t.Name, &t.Status, &t.EndDate)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	err = row.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *TaskModel) Find(id int) (models.Tasks, error) {
	var t models.Tasks

	stmt := `SELECT id,name,status,endDate FROM tasks WHERE ID = ?`
	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&t.ID, &t.Name, &t.Status, &t.EndDate)
	if err != nil {
		return models.Tasks{}, err
	}

	return t, nil
}

func (m *TaskModel) Like(search string) ([]models.Tasks, error) {
	stmt := fmt.Sprintf("SELECT id,name,status,endDate FROM tasks WHERE name LIKE %q", "%"+search+"%")

	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	tasks := []models.Tasks{}
	for row.Next() {
		t := models.Tasks{}
		err := row.Scan(&t.ID, &t.Name, &t.Status, &t.EndDate)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	err = row.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *TaskModel) Insert(name, status, date string) error {
	stmt := `INSERT INTO tasks (name,status,endDate) VALUES (?,?,?)`

	_, err := m.DB.Exec(stmt, name, status, date)
	return err
}

func (m *TaskModel) Delete(id int) error {
	stmt := `DELETE FROM tasks WHERE id=?`
	_, err := m.DB.Exec(stmt, id)

	return err
}

func (m *TaskModel) Update(name, status, date string, id int) error {
	stmt := `UPDATE tasks SET name = $1,status = $2,endDate = $3 WHERE id = $4`
	_, err := m.DB.Exec(stmt, name, status, date, id)

	return err
}
