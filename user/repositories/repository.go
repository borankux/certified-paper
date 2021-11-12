package repositories

import (
	"errors"
	"github.com/borankux/certified-paper/dto"
)


func FindUserById(id int) (dto.User, error){
	u := dto.User{
		ID:       1,
		Name:     "",
		Password: "",
		Contacts: []dto.UserContact{
			{
				ID:    1,
				Value: "mirzatsoft@gmail.com",
				Type:  dto.ContactTypeEmail,
			},
		},
	}

	if u.ID == id {
		return u, nil
	}

	return dto.User{}, errors.New("cannot find user")
}