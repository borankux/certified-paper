package database

import (
	"fmt"
	"github.com/borankux/certified-paper/dto"
)

type FileDB struct {
	Database
}

var fileName string

func init() {
	fileName := "out/database.txt"
	fmt.Println("Filename" + fileName)
}

func (FileDB) Init() error {
	fmt.Println(fileName)
	return nil
}

func (FileDB) DeleteUser(id int) error {
	fmt.Println("Removed a user")
	return nil
}

func (FileDB) InsertUser(user dto.User) error {
	fmt.Printf("inserting a user %v\n", user)
	return nil
}

func(FileDB) GetUserById(id int) dto.User {
	return dto.User{}
}

func (FileDB) UserExists(email string) error {
	return nil
}