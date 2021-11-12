package database

import (
	"fmt"
)

type FileDB struct {
	Database
}

var fileName string

func init() {
	fileName := "out/database.txt"
	fmt.Println("Filename" + fileName)
}

func (FileDB) Init() bool {
	fmt.Println(fileName)
	return true
}
