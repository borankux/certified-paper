package user

import (
	"github.com/borankux/certified-paper/database"
	"github.com/borankux/certified-paper/dto"
)

var db database.Database

func init() {
	db = database.Sqlite{}
	err := db.Init()
	if err != nil {
		return 
	}
}


func Register(email string, password string) {
	err := db.InsertUser(dto.User{
		ID:       0,
		Name:     "",
		Email:    email,
		Password: password,
		Contacts: nil,
	})

	if err != nil {
		//print log here
	}
}
