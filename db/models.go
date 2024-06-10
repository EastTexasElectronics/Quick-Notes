package db

import (
    "log"
    "time"
)

type Note struct {
    ID        int
    Title     string
    Content   string
    CreatedAt time.Time
}

func CreateNote(title, content string) {
    _, err := DB.Exec("INSERT INTO notes (title, content) VALUES (?, ?)", title, content)
    if err != nil {
        log.Fatal("Cannot create note:", err)
    }
}

func GetNotes() ([]Note, error) {
    rows, err := DB.Query("SELECT id, title, content, created_at FROM notes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notes []Note
    for rows.Next() {
        var note Note
        err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
        if err != nil {
            return nil, err
        }
        notes = append(notes, note)
    }
    return notes, nil
}

func UpdateNote(id int, title, content string) {
    _, err := DB.Exec("UPDATE notes SET title = ?, content = ? WHERE id = ?", title, content, id)
    if err != nil {
        log.Fatal("Cannot update note:", err)
    }
}

func DeleteNote(id int) {
    _, err := DB.Exec("DELETE FROM notes WHERE id = ?", id)
    if err != nil {
        log.Fatal("Cannot delete note:", err)
    }
}
