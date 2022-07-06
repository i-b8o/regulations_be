package dto

// CreateParagraphsParams used by CreateParagraphs.
type CreateParagraphsParams struct {
	Paragraphs []Paragraph
}

type Paragraph struct {
	ID                string
	Href              string
	ParagraphOrderNum string
	ParagraphClass    string
	ParagraphText     string
	ChapterID         uint64
}
