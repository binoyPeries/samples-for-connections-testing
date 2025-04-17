-- Create a new database
CREATE DATABASE IF NOT EXISTS users_db;

-- Use the new database
USE users_db;

-- Create a table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

-- Insert sample data
INSERT INTO users (name, email) VALUES ('Alice', 'alice@example.com');
INSERT INTO users (name, email) VALUES ('Bob', 'bob@example.com');

-- Create a new user and grant privileges
CREATE USER IF NOT EXISTS 'sample_user'@'%' IDENTIFIED BY 'sample_pass';
GRANT ALL PRIVILEGES ON users_db.* TO 'sample_user'@'%';
FLUSH PRIVILEGES;
