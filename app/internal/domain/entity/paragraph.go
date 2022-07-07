package entity

type Paragraph struct {
	ID        string
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
