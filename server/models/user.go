package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email    string `json:"email" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
    Notes    []Note `json:"notes" gorm:"foreignKey:UserID"`
}

// Create a new user in the database
func CreateUser(user *User) error {
    result := DB.Create(user)
    return result.Error
}

// Get a user by email
func GetUserByEmail(user *User, email string) error {
    result := DB.Where("email = ?", email).First(user)
    return result.Error
}

// Get a user by ID
func GetUserByID(user *User, id uint) error {
    result := DB.First(user, id)
    return result.Error
}
