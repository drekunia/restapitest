package main

import (
	"time"
)

type User struct {
	ID           uint       `json:"id" sql:"primary key;not null; unique"`
	Username     string     `json:"username" sql:"not null; unique"`
	Email        string     `json:"email" sql:"not null; unique"`
	PasswordHash string     `json:"password_hash" sql:"not null"`
	FirstName    string     `json:"first_name" sql:"not null"`
	LastName     string     `json:"last_name" sql:"not null"`
	City         string     `json:"city"`
	Country      string     `json:"country"`
	Avatar       string     `json:"avatar"`
	Bio          string     `json:"bio"`
	CreatedAt    time.Time  `json:"created_at" sql:"not null"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
}
