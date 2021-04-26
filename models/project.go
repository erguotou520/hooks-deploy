package models

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Type        string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Token       string
	TestScript  string
	BuildScript string `gorm:"default:yarn build"`
	DistDir     string `gorm:"default:dist"`
	SPA         bool   `gorm:"default:true"`
	Branch      string
	Tag         bool
	TagMatcher  string
	Domain      string
	Rewrites    datatypes.JSON
}

type Rewrite struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ProjectForm struct {
	ID          uint      `json:"id" form:"id"`
	Type        string    `json:"Type" form:"Type"`
	Name        string    `json:"name" form:"name" binding:"required"`
	Token       string    `json:"token" form:"token"`
	TestScript  string    `json:"testScript" form:"testScript"`
	BuildScript string    `json:"buildScript" form:"buildScript"`
	DistDir     string    `json:"distDir" form:"distDir"`
	SPA         bool      `json:"spa" form:"spa"`
	Branch      string    `json:"branch" form:"branch"`
	Tag         bool      `json:"tag" form:"tag"`
	TagMatcher  string    `json:"tagMatcher" form:"tagMatcher"`
	Domain      string    `json:"domain" form:"domain"`
	Rewrites    []Rewrite `json:"rewrites" form:"rewrites"`
}

func (f ProjectForm) ToProject() Project {
	project := Project{
		Type:        f.Type,
		Name:        f.Name,
		Token:       f.Token,
		TestScript:  f.TestScript,
		BuildScript: f.BuildScript,
		DistDir:     f.DistDir,
		SPA:         f.SPA,
		Branch:      f.Branch,
		Tag:         f.Tag,
		TagMatcher:  f.TagMatcher,
		Domain:      f.Domain,
	}
	rewrites, err := json.Marshal(f.Rewrites)
	if err == nil && len(rewrites) > 0 {
		project.Rewrites = rewrites
	}
	return project
}

func (p Project) ToProjectForm() ProjectForm {
	project := ProjectForm{
		Type:        p.Type,
		Name:        p.Name,
		Token:       p.Token,
		TestScript:  p.TestScript,
		BuildScript: p.BuildScript,
		DistDir:     p.DistDir,
		SPA:         p.SPA,
		Branch:      p.Branch,
		Tag:         p.Tag,
		TagMatcher:  p.TagMatcher,
		Domain:      p.Domain,
	}
	// if p.Rewrites != nil {
	// 	var rewritesJSON []Rewrite
	// 	rewrites, err := json.Unmarshal([]byte(p.Rewrites.String()), &rewritesJSON)
	// 	if err != nil {
	// 		project.Rewrites = rewrites
	// 	}

	// }
	return project
}
