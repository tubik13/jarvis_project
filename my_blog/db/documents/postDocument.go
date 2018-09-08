package documents

type PostDokument struct {
	ID              string `bson:"id, omit_empty"`
	Title           string
	ContentHtml     string
	ContentMarkdown string
}
