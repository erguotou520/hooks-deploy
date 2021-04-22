package main

import (
	"codeup-auto-deploy/admin"
	"codeup-auto-deploy/db"
	"codeup-auto-deploy/webhooks"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gin.DisableConsoleColor()
	r := gin.New()

	// logger
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())

	// recovery
	r.Use(gin.Recovery())

	// templates
	r.LoadHTMLGlob("admin/html/*.html")

	db.InitDB()

	// routes
	webhookG := r.Group("/webhook")
	webhookG.POST("/codeup", webhooks.CodeupHooks)

	// users
	adminUser := "admin"
	adminPassword := "password"
	if value, existed := os.LookupEnv("ADMIN_USER"); existed {
		adminUser = value
	}
	if value, existed := os.LookupEnv("ADMIN_PASSWORD"); existed {
		adminPassword = value
	}
	var adminAccounts = make(map[string]string)
	adminAccounts[adminUser] = adminPassword
	log.Println(adminAccounts)
	adminG := r.Group("/admin", gin.BasicAuth(adminAccounts))
	{
		adminG.GET("/", admin.AdminPage)
		adminG.GET("/new", admin.FormPage)
		adminG.GET("/edit/:id", admin.FormPage)
		adminG.POST("/save", admin.SaveForm)
	}

	// test api
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
