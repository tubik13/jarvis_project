package models

import "html/template"

type Post struct {
	Id              string
	Title           string
	ContentHtml     string
	ContentMarkdown string
}

func NewPost(id, title string, contentHtml template.HTML, contentMarkdown string) *Post {
	return &Post{id, title, string(contentHtml), contentMarkdown}
}
