package dto_test

import (
	"github.com/borankux/certified-paper/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserStruct(t *testing.T) {
	contacts := []dto.UserContact{
		{
			ID:    0,
			Value: "18621577147",
			Type:  dto.ContactTypePhone,
		},
		{
			ID:    1,
			Value: "test@test.com",
			Type:  dto.ContactTypeEmail,
		},
	}
	testUser := dto.User{
		ID:       0,
		Name:     "test_user",
		Password: "test_password",
		Contacts: contacts,
	}
	assert.Equal(t, "test_user", testUser.Name)
	assert.Equal(t, "test_password", testUser.Password)
	assert.Equal(t, dto.ContactTypeEmail, testUser.Contacts[1].Type)
	assert.Equal(t, "test@test.com", testUser.Contacts[1].Value)
}
