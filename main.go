package main

import (
	"GOTODO/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

var db *sql.DB

// TODO: Add Middleware
// TODO: Generate better IDs
// TODO: Handle passwords securely

func main() {

	var err error
	db, err = sql.Open("sqlite", "./db/development.db")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", StartHandler).Methods("GET")

	/* r.Use(loggingMiddleware) */

	fmt.Print("Server started on port 8080\n")

	// 1. Create todo (title + task) for a specific user
	// * if user does not exist, create a new user with the given username
	// * if user exists, use the existing user by getting the user ID
	// * if username is not provided, return an error
	// * if title or task is not provided, return an error
	// * if title and task are provided, create a new todo
	// * if task already exists with the same title, return an error
	r.HandleFunc("/user/{username}/todo/", createTodoHandler).Methods("POST")

	// 3. Edit/update task on a specific todo - title and completed state
	r.HandleFunc("/user/{username}/todo/", updateTodoHandler).Methods("PUT")

	// 4. Delete task on a specific todo - title, tasks, completed state
	r.HandleFunc("/user/{username}/todo/", deleteTaskHandler).Methods("DELETE")

	// 5. Get all todo lists (title) for a specific user
	r.HandleFunc("/user/{username}/title/", getAllTitlesHandler).Methods("GET")

	// 6. Get all tasks for a specific todo - title
	//r.HandleFunc("/user/{username}/todo/{title}/tasks", getAllTasksHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the TODO API!\n")
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	if todo.Title == "" || todo.Task == "" {
		http.Error(w, "Title and task are required", http.StatusBadRequest)
		return
	}

	// * Does user already exist?
	// * If not, create a new user with the given username
	userID, err := getOrCreateUser(username)
	if err != nil {
		http.Error(w, "Failed to get or create user in createTodoHandler", http.StatusInternalServerError)
		return
	}

	// Insert todo in db for the user
	// * A user can not have multiple todos with the same title and task (database constraint)
	_, err = db.Exec("INSERT INTO todos (user_id, title, task) VALUES (?, ?, ?)", userID, todo.Title, todo.Task)
	if err != nil {
		http.Error(w, "Failed to save todo. Todo might already exist", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Created todo for user '%s': [%s] %s\n", username, todo.Title, todo.Task)
}

// TODO: Handle passwords
// TEST: This function is used in multiple handlers, and might create a handler where it should not - e.g. delete handler? what happens?
func getOrCreateUser(username string) (int64, error) {
	var userID int64
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)

	if err == sql.ErrNoRows {
		// User does not exist, create a new user
		res, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, "defaultPassword")
		if err != nil {
			return 0, err
		}
		userID, err = res.LastInsertId()
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	return userID, nil
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	if req.OldTitle == "" || req.Task == "" || req.NewTitle == "" {
		http.Error(w, "old_title, new_title, and task are required", http.StatusBadRequest)
		return
	}

	userID, err := getOrCreateUser(username)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(
		"UPDATE todos SET title = ?, completed = ? WHERE user_id = ? AND title = ? AND task = ?",
		req.NewTitle, req.Completed, userID, req.OldTitle, req.Task,
	)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Updated todo for user '%s': [%s] %s -> [%s], completed: %v\n", username, req.OldTitle, req.Task, req.NewTitle, req.Completed)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	var req models.Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Task == "" {
		http.Error(w, "title and task are required", http.StatusBadRequest)
		return
	}

	userID, err := getOrCreateUser(username)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	res, err := db.Exec(
		"DELETE FROM todos WHERE user_id = ? AND title = ? AND task = ?",
		userID, req.Title, req.Task,
	)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	// check if any rows were affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check delete result", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Deleted todo for user '%s': [%s] %s\n", username, req.Title, req.Task)
}

func getAllTitlesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	userID, err := getOrCreateUser(username)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT title FROM todos WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var titles []string
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			http.Error(w, "Failed to scan todo title", http.StatusInternalServerError)
			return
		}
		titles = append(titles, title)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading todos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(titles)
}
