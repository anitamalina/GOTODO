CREATE TABLE IF NOT EXISTS `users` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `username` TEXT NOT NULL UNIQUE,
  `password` TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `todos`(
  user_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  task TEXT NOT NULL,
  completed BOOLEAN NOT NULL DEFAULT 0,
  PRIMARY KEY (user_id, title, task),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert sample data for James Cooper --
INSERT INTO `users` (username, password) VALUES ('JamesCooper', '123James');
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Grocery Shopping', 'Coffee', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Grocery Shopping', 'Milk', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Grocery Shopping', 'Cheese', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Movie Night', 'Inception', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Movie Night', 'Interstellar', 1);

-- Insert sample data for Alice Smith --
INSERT INTO `users` (username, password) VALUES ('AliceSmith', '123Alice');
INSERT INTO `todos` (user_id, title, task, completed) VALUES (2, 'Book Club', 'Under the Rainbow', 0);

/* DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `todos`; */

/* SELECT * FROM users; */
/* DELETE FROM users WHERE username = 'AliceSmith'; */

