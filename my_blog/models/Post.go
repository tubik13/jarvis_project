package models

import "html/template"

type Post struct {
	Id              string
	Title           string
	ContentHtml     template.HTML
	ContentMarkdown string
}

func NewPost(id, title string, contentHtml template.HTML, contentMarkdown string) *Post {
	return &Post{id, title, contentHtml, contentMarkdown}
}
