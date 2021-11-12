package requests_test

import (
	"github.com/borankux/paperstamp/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterRequestStruct(t *testing.T) {
	req := requests.RegisterRequest{
		Email: "mablat@ebay.com",
		Password: "123456",
	}
	assert.Equal(t, req.Email, "mablat@ebay.com")
	assert.Equal(t, req.Password, "123456")
}
