-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    endDate TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
