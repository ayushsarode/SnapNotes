package models

import (
    "gorm.io/gorm"

)

type Note struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}

func CreateNote(note *Note) error {
    result := DB.Create(note)
    return result.Error
}

func GetNotesByUser(userID uint) ([]Note, error) {
    var notes []Note
    result := DB.Where("user_id = ?", userID).Find(&notes)
    return notes, result.Error
}

func UpdateNoteByID(userID uint, noteID string, note *Note) (*Note, error) {
    var existingNote Note
    result := DB.Where("id = ? AND user_id = ?", noteID, userID).First(&existingNote)
    if result.Error != nil {
        return nil, result.Error
    }

    existingNote.Title = note.Title
    existingNote.Content = note.Content
    DB.Save(&existingNote)

    return &existingNote, nil
}

func DeleteNoteByID(userID uint, noteID string) error {
    result := DB.Where("id = ? AND user_id = ?", noteID, userID).Delete(&Note{})
    return result.Error
}
