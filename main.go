package main

import (
	"github.com/borankux/paperstamp/admin"
	"github.com/borankux/paperstamp/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	//userRoutes
	userRoutes := app.Group("/userRoutes")
	userRoutes.POST("/login", user.Login)
	userRoutes.POST("/logout", user.Logout)
	userRoutes.POST("/register", user.Register)
	userRoutes.POST("/code", user.SendCode)
	userRoutes.POST("/articles", user.AddArticle)
	userRoutes.DELETE("/articles", user.DeleteArticle)
	userRoutes.GET("/articles:id", user.GetArticles)
	userRoutes.GET("/articles", user.GetArticle)

	//adminRoutes
	adminRoutes := app.Group("/admin")
	adminRoutes.POST("/login", admin.Login)
	adminRoutes.POST("/logout", admin.Logout)
	adminRoutes.GET("/articles", admin.GetArticles)
	adminRoutes.GET("/articles:id", admin.GetArticle)
	adminRoutes.PUT("/articles", admin.UpdateArticle)
	adminRoutes.GET("/users:id", admin.GetUser)
	adminRoutes.GET("/users", admin.GetUsers)
	adminRoutes.GET("/crawlers", admin.GetCrawlers)
	adminRoutes.GET("/sys", admin.GetSysInfo)
	adminRoutes.GET("/search", admin.Search)

	err := app.Run(":8080")
	if err != nil {
		return
	}
}