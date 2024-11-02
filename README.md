## How to run

You need to have installed:

* Go 1.23+
* Sqlite3
* [Goose](https://github.com/pressly/goose ) (optional,but recommended)

At the root folder create a new sqlite instance named app.db

```
goose --dir=assets/migrations sqlite3 app.db up
```

Then run the application with

```
go run cmd/web/*
```
