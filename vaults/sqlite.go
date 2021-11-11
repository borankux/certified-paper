package vaults

import (
	"github.com/fatih/color"
)

type SQLite struct {
	Vault
}

func (SQLite) Save(target interface{}) bool{
	color.Green("Saving stuff")
	return true
}

func (SQLite) Update(id int, target interface{}) bool{
	color.Yellow("Updating stuff: %d", id)
	return true
}

func (SQLite) Delete(id int) bool {
	color.Red("Deleting stuff: %d ", id)
	return true
}

func (SQLite) Find(id int, filter interface{}) interface{}{
	color.Blue("Finding Stuff: %d ", id)
	return nil
}

func (SQLite) Search(filter interface{}) interface{} {
	return nil
}