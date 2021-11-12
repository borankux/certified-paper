package database

import (
	"github.com/borankux/certified-paper/dto"
)

type Database interface {
	Init() bool
	InsertUser(user dto.User)
	GetUserById(id int) dto.User
	UserExists(email string) bool
	DeleteUser(id int) bool
}