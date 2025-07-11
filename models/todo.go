package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Todo struct {
	UserID    int64  `json:"user_id"`
	Title     string `json:"title"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type UpdateTodoRequest struct {
	OldTitle  string `json:"old_title"`
	Task      string `json:"task"`
	NewTitle  string `json:"new_title"`
	Completed bool   `json:"completed"`
}
