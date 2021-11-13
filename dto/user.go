package dto

const (
	ContactTypeEmail = iota
	ContactTypePhone
)

type User struct {
	ID int
	Name string
	Email string
	Password string
	Contacts []UserContact
}

type UserContact struct {
	ID int
	Value string
	Type int
}
