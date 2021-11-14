package database

import (
	"database/sql"
	"fmt"
)

type Sqlite struct {
	Database
}

func (Sqlite) Init() error {
	_, err := sql.Open("sqlite", "data/database.db")
	if err != nil {
		fmt.Println("have some error")
	}
	return nil
}