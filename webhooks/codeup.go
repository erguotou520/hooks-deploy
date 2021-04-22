package webhooks

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type CodeupHooksData struct {
	ObjectKind string `json:"object_kind"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Ref        string `json:"ref"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserEmail  string `json:"user_email"`
	ProjectID  int    `json:"project_id"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}

func CodeupHooks(c *gin.Context) {
	event := c.GetHeader("X-Codeup-Event")
	if event == "Push Hook" {
		var hookData CodeupHooksData
		err := c.BindJSON(&hookData)
		if err == nil {
			log.Println(111)
			log.Println(hookData.Repository.Name)
			log.Println(hookData.Repository.GitSSHURL)
		}
	}
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
