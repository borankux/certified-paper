package database

import "github.com/borankux/certified-paper/dto"

type Database interface {
	Init() error
	InsertUser(user dto.User) error
	GetUserById(id int) dto.User
	UserExists(email string) error
	DeleteUser(id int) error
}