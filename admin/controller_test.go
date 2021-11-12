package admin_test

import (
	"github.com/borankux/certified-paper/admin"
	"github.com/gin-gonic/gin"
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

func TestLogin(t *testing.T) {
	admin.Login(ctx)
}

func TestLogout(t *testing.T) {
	admin.Logout(ctx)
}

func TestGetUser(t *testing.T) {
	admin.GetUser(ctx)
}


func TestGetUsers(t *testing.T) {
	admin.GetUsers(ctx)
}

func TestGetArticle(t *testing.T) {
	admin.GetArticle(ctx)
}

func TestGetArticles(t *testing.T) {
	admin.GetArticles(ctx)
}

func TestGetCrawlers(t *testing.T) {
	admin.GetCrawlers(ctx)
}


func TestGetSysInfo(t *testing.T) {
	admin.GetSysInfo(ctx)
}

func TestUpdateArticle(t *testing.T) {
	admin.UpdateArticle(ctx)
}

func TestSearch(t *testing.T) {
	admin.Search(ctx)
}