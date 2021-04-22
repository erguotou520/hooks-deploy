package admin

import (
	"codeup-auto-deploy/db"
	"codeup-auto-deploy/models"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminPage(c *gin.Context) {
	log.Println("home")
	var projects []models.Project
	db.DB.Find(&projects)
	log.Println(projects)
	c.HTML(http.StatusOK, "home.html", gin.H{
		"projects": projects,
	})
}

func FormPage(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	if id != "" {
		result := db.DB.First(&project, id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.HTML(http.StatusOK, "404.html", nil)
		} else {
			c.HTML(http.StatusOK, "form.html", gin.H{
				"project": project,
			})
		}
	} else {
		c.HTML(http.StatusOK, "form.html", nil)
	}
}

func SaveForm(c *gin.Context) {
	var project models.Project
	err := c.Bind(&project)
	if err == nil {
		if project.ID == 0 {
			db.DB.Create(&project)
		} else {
			db.DB.Save(&project)
		}
		c.HTML(http.StatusOK, "saved.html", gin.H{
			"id": project.ID,
		})
	} else {
		log.Println(err)
	}
}
