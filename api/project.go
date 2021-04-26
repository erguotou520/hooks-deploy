package api

import (
	"codeup-auto-deploy/db"
	"codeup-auto-deploy/models"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project
	db.DB.Find(&projects)
	c.JSON(http.StatusOK, projects)
}

func GetProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	result := db.DB.First(&project, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

func SaveProject(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{
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
