
Create DATABASE notes_app;

USE notes_app;

CREATE TABLE notes (
    id INT AUTO_INCREMENT PRIMARY KEY, -- Primary key
    title VARCHAR(255) NOT NULL, -- Title of the note
    content TEXT NOT NULL, -- Content of the note
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Creation timestamp
);