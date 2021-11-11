package atoms

const (
	EMAIL = iota
	PHONE
)

type User struct {
	ID int
	Name string
	Password string
}

type UserContact struct {
	ID int
	Value string
	Type int
}
