package repositories_test

import (
	"github.com/borankux/paperstamp/user/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUserById__SUCCESS(t *testing.T) {
	_, err := repositories.FindUserById(1)
	assert.Nil(t, err)
}


func TestFindUserById__FAIL(t *testing.T) {
	_, err := repositories.FindUserById(2)
	assert.NotNil(t, err)
}
