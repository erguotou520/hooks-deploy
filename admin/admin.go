package admin

import (
	"codeup-auto-deploy/db"
	"codeup-auto-deploy/models"
	"errors"
	"log"
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminPage(c *gin.Context) {
	var projects []models.Project
	db.DB.Find(&projects)
	ginview.HTML(c, http.StatusOK, "home", gin.H{
		"projects": projects,
	})
}

func FormPage(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		var project models.ProjectForm
		result := db.DB.First(&project, id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ginview.HTML(c, http.StatusOK, "404", nil)
		} else {
			ginview.HTML(c, http.StatusOK, "form", gin.H{
				"project": project,
			})
		}
	} else {
		ginview.HTML(c, http.StatusOK, "form", gin.H{
			"project": models.Project{
				Type:    "codeup",
				DistDir: "dist",
				SPA:     true,
			},
		})
	}
}

func SaveForm(c *gin.Context) {
	var project models.ProjectForm
	err := c.Bind(&project)
	if err == nil {
		log.Println(project)
		var dbProject = project.ToProject()
		if project.ID == 0 {
			db.DB.Create(&dbProject)
		} else {
			db.DB.Save(&dbProject)
		}
		ginview.HTML(c, http.StatusOK, "saved", gin.H{
			"id": dbProject.ID,
		})
	} else {
		log.Println(err)
	}
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	db.DB.Delete(&models.Project{}, id)
	c.Status(http.StatusOK)
}
