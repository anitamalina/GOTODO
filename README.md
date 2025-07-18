# Go TODO API

A simple RESTful TODO list API written in Go using SQLite and Gorilla Mux.

## Features

- Create todos (title + task) for specific users
- Update existing todos (change title and mark as completed)
- Delete tasks
- List all todo titles for a user

## Requirements

- Go 1.20 or newer
- SQLite (VSCode extension)

## Setup

1. **Clone the repository**
    ```bash
    git clone https://github.com/anitamalina/GOTODO.git
    cd GOTODO
    ```

2. **Create the SQLite database**

    Ensure you have a database file at 
    `./db/development.db`

    In VSCode open search and write: 
    ```text
    >SQLite: Open Database
    ```

    and select the database `./db/development.db`
    
    `SQLite Explorer` will open in left lower corner, where the database can be refreshed and tables be "played" to see it visually.

    You might have to run the queries in `./db/sql.sql`
    
    In VSCode open search and write: 
    ```text
    >SQLite: Run Query
    ```

3. **Run the server**
    ```bash
    go run main.go
    ```

    You should see:
    ```text
    Server started on port 8080
    ```

## Example Request

Create to do for new user:
http://localhost:8080/user/AndreaSimson/todo/

```json
{
  "title": "Travels",
  "task": "Dolomites"
}
```

Create new to do list for exsisting user:
http://localhost:8080/user/JamesCooper/todo/

```json
{
  "title": "Buy LEGO",
  "task": "Titanic"
}
```

Add task to an exsisting todo list for exsisting user:
http://localhost:8080/user/JamesCooper/todo/

```json
{
  "title": "Buy LEGO",
  "task": "Star wars"
}
```

Update to do title for AndreaSimson:
http://localhost:8080/user/AndreaSimson/todo/

```json
{
  "old_title": "Travels",
  "new_title": "Places I want to go",
	"task": "Dolomites",
  "completed": false
}
```

Update todo completion for JamesCooper:
http://localhost:8080/user/JamesCooper/todo/
```json
{
  "old_title": "Movie Night",
  "new_title": "Movie Night",
	"task": "Inception",
  "completed": true
}
```

Delete todo from AndreaSimson:
http://localhost:8080/user/AndreaSimson/todo/

```json
{
  "title": "Places I want to go",
  "task": "Dolomites"
}
```

Get todo titles for AliceSmith: 
http://localhost:8080/user/AliceSmith/


Get tasks for a todo title for JamesCooper:
http://localhost:8080/user/JamesCooper/Grocery%20Shopping/tasks

