package main_test

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	exitFlag := m.Run()
	os.Exit(exitFlag)
}