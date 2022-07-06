package entity

type Paragraph struct {
	ID        string
	Href      int
	Num       string
	Class     string
	Content   string
	ChapterID uint64
}

type ParagraphsInput struct {
	Paragraphs []Paragraph
}

type ParagraphsOutput struct {
	Message string
}
