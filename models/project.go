package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name" form:"name" binding:"required" gorm:"not null"`
	Token       string `json:"token" form:"token"`
	TestScript  string `json:"testScript" form:"testScript"`
	BuildScript string `json:"buildScript" form:"buildScript"`
	DistDir     string `json:"distDir" form:"distDir" binding:"required" gorm:"not null"`
	Branch      string `json:"branch" form:"branch"`
	Tag         string `json:"tag" form:"tag"`
	TagMatcher  string `json:"tagMatcher" form:"tagMatcher"`
	Domain      string `json:"domain" form:"domain" binding:"required" gorm:"not null"`
	Rewrites    string `json:"rewrites" form:"rewrites"`
}

type ProjectForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Token       string `json:"token"`
	TestScript  string `json:"testScript"`
	BuildScript string `json:"buildScript"`
	DistDir     string `json:"distDir" binding:"required"`
	Branch      string `json:"branch"`
	Tag         string `json:"tag"`
	TagMatcher  string `json:"tagMatcher"`
	Domain      string `json:"domain" binding:"required"`
	Rewrites    []struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"rewrites"`
}
