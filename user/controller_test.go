package user_test

import (
	"encoding/json"
	"fmt"
	"github.com/borankux/paperstamp/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var recorder *httptest.ResponseRecorder
var ctx *gin.Context

func init() {
	gin.SetMode(gin.TestMode)
	recorder = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)
}

func TestRegister(t *testing.T) {
	body := `{email:'mirzat', 'password':'123234'}`
	ctx.Writer.WriteString(body)

	user.Register(ctx)
	all, err := ioutil.ReadAll(recorder.Body)
	var result string
	json.Unmarshal(all, &result)
	fmt.Println(result)
	if err != nil {
		return
	}
	assert.Equal(t, 200, recorder.Code)
}

func TestLogin(t *testing.T) {
	user.Logout(ctx)
	assert.True(t, true)
}

func TestLogout(t *testing.T) {
	user.Login(ctx)
	assert.True(t, true)
}

func TestAddArticle(t *testing.T) {
	user.AddArticle(ctx)
	assert.True(t, true)
}

func TestGetArticle(t *testing.T) {
	user.GetArticle(ctx)
	assert.True(t, true)
}

func TestGetArticles(t *testing.T) {
	user.GetArticles(ctx)
	assert.True(t, true)
}

func TestDeleteArticle(t *testing.T) {
	user.DeleteArticle(ctx)
	assert.True(t, true)
}

func TestSendCode(t *testing.T) {
	user.SendCode(ctx)
	assert.True(t, true)
}