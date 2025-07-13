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
INSERT INTO `users` (username, password) VALUES ('JamesCooper', '$2a$14$X6r9xS/2W5IF5kV9Bq.JruZIQ98t.2g0BfD5JX2R2BuO2W93W5sI2');
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Grocery Shopping', 'Coffee', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Grocery Shopping', 'Milk', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Grocery Shopping', 'Cheese', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Movie Night', 'Inception', 0);
INSERT INTO `todos` (user_id, title, task, completed) VALUES (1, 'Movie Night', 'Interstellar', 1);

-- Insert sample data for Alice Smith --
INSERT INTO `users` (username, password) VALUES ('AliceSmith', '$2a$14$wD1c2u9u4F9KM3s9e9NqWOD3AtZKXmbW6WEmFG3Ai6mKJ0Z7o.Eji');
INSERT INTO `todos` (user_id, title, task, completed) VALUES (2, 'Book Club', 'Under the Rainbow', 0);
 
/* DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `todos`; */

/* DELETE FROM users WHERE username = 'AndreaSimson'; */

