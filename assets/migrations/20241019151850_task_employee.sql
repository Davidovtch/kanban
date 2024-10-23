-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_employee(
    id INT PRIMARY KEY AUTOINCREMENT,    
    task_id INT NOT NULL,
    employee_id INT NOT NULL,
    FOREIGN KEY(task_id) REFERENCES tasks(id),
    FOREIGN KEY(employee_id) REFERENCES employees(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_employee;
-- +goose StatementEnd
