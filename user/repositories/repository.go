package user

import (
	"errors"
	"github.com/borankux/paperstamp/dto"
)


func FindUserById(id int) (interface{}, error){
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

	return nil, errors.New("cannot find user")
}