package dto

const (
	CONTACT_TYPE_EMAIL = iota
	CONTACT_TYPE_PHONE
)

type User struct {
	ID int
	Name string
	Password string
	Contacts []UserContact
}

type UserContact struct {
	ID int
	Value string
	Type int
}
