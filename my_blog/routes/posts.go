package routes

import (
	"my_blog/db/documents"
	"my_blog/models"
	"my_blog/session"
	"my_blog/utils"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func WriteHandler(rnd render.Render, s *session.Session) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	model := models.EditPostModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Post = models.Post{}
	rnd.HTML(200, "write", model)
}
func EditHandler(rnd render.Render, s *session.Session, db *mgo.Database, params martini.Params) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	postsCollection := db.C("posts")

	id := "id"

	//fmt.Println(params)

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	model := models.EditPostModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Post = post
	rnd.HTML(200, "write", model)
}

func ViewHandler(s *session.Session, rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {
	postsCollection := db.C("posts")

	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	model := models.ViewPostModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Post = post
	rnd.HTML(200, "view", model)
}

func SavePostHandler(s *session.Session, rnd render.Render, r *http.Request, db *mgo.Database) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownToHtml(contentMarkdown)

	postDocument := documents.PostDocument{id, title, contentHtml, contentMarkdown}

	postsCollection := db.C("posts")
	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = utils.GenerateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	rnd.Redirect("/")
}

func DeleteHandler(s session.Session, rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}
	postsCollection := db.C("Posts")
	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}

func GetHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkdownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}
