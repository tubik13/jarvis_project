package routes

import (
	"fmt"
	"my_blog/db/documents"
	"my_blog/models"
	"my_blog/session"

	"gopkg.in/mgo.v2"

	"github.com/martini-contrib/render"
)

func IndexHandler(rnd render.Render, s *session.Session, db *mgo.Database) {
	fmt.Println(s.Username)

	postDocuments := []documents.PostDocument{}
	postsCollection := db.C("posts")
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	model := models.PostListModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Posts = posts

	rnd.HTML(200, "index", model)
}
