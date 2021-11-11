package vaults

type Vault interface {
	Save(target interface{}) bool
	Update(id int, target interface{}) bool
	Delete(id int) bool
	Find(id int, filters interface{}) interface{}
	Search(target interface{}) interface{}
}