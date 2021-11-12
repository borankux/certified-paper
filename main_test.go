package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	exitFlag := m.Run()
	os.Exit(exitFlag)
}

func TestRunMain(t *testing.T) {
	go main()
	res, _ := http.Get("http://localhost:8080")
	assert.Equal(t, 200, res.StatusCode)
}

func TestRunMain_failed(t *testing.T) {
	go main()
	go main()
}