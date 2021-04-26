package main

import (
	"codeup-auto-deploy/api"
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
	// r.LoadHTMLGlob("admin/html/*.html")
	// r.HTMLRender = ginview.New(goview.Config{
	// 	Root:      "views/admin",
	// 	Extension: ".html",
	// 	Master:    "layouts/master",
	// 	Partials:  []string{},
	// 	// Partials:  []string{"partials/ad"},
	// 	// Funcs: template.FuncMap{
	// 	// 	"copy": func() string {
	// 	// 		return time.Now().Format("2006")
	// 	// 	},
	// 	// },
	// 	DisableCache: true,
	// })

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

	// admin template
	// adminTemp := ginview.NewMiddleware(goview.Config{
	// 	Root:      "views/admin",
	// 	Extension: ".html",
	// 	Master:    "layouts/master",
	// 	Partials:  []string{},
	// 	// Partials:  []string{"partials/ad"},
	// 	// Funcs: template.FuncMap{
	// 	// 	"copy": func() string {
	// 	// 		return time.Now().Format("2006")
	// 	// 	},
	// 	// },
	// 	DisableCache: true,
	// })

	// adminG := r.Group("/admin", gin.BasicAuth(adminAccounts), adminTemp)
	// {
	// 	adminG.GET("/", admin.AdminPage)
	// 	adminG.GET("/new", admin.FormPage)
	// 	adminG.GET("/edit/:id", admin.FormPage)
	// 	adminG.POST("/save", admin.SaveForm)
	// 	adminG.DELETE("/project/:id", admin.DeleteProject)
	// }
	apiG := r.Group("/api", gin.BasicAuth(adminAccounts))
	{
		apiG.GET("/projects", api.GetProjects)
		apiG.GET("/projects/:id", api.GetProject)
		apiG.POST("/projects", api.SaveProject)
		apiG.DELETE("/project/:id", api.DeleteProject)
	}

	// test api
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
