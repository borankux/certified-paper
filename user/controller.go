package user

import (
	"fmt"
	"github.com/borankux/paperstamp/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	//will do
}

func Logout(ctx *gin.Context) {
	//will do
}

func Register(ctx *gin.Context) {
	req := requests.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)

	//validate required
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(req)
}

func SendCode(ctx *gin.Context) {
	//will do
}

func AddArticle(ctx *gin.Context) {
	//will do
}

func DeleteArticle(ctx *gin.Context) {
	//will do
}

func GetArticles(ctx *gin.Context) {
	//will do
}

func GetArticle(ctx *gin.Context) {
	//will do
}