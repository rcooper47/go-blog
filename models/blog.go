package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string
	Slug  string
	Body  string
	/*
		Title string
		Slug string
		Body string		Use body bc then you can already have the page layout handled.  Also gives i a chance to practice with HTML.
		SubTitle string
		HeaderImg string
		ingredients
		Description







		title
		subtitle
		body
		ingredients
		tags
		tableOfContents / subTitles




	*/
}

type BlogRequest struct {
	Title string
	Slug  string
	Body  string
}
