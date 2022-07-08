package entity

type Paragraph struct {
	ID        uint64 `json:"paragraph_id"`
	Num       uint64 `json:"paragraph_num"`
	IsHTML    bool   `json:"paragraph_ishtml"`
	Class     string `json:"paragraph_class,omitempty"`
	Content   string `json:"paragraph_content,omitempty"`
	ChapterID uint64 `json:"paragraph_chapterid,omitempty"`
}

type ParagraphsInput struct {
	Paragraphs []Paragraph
}

type ParagraphsOutput struct {
	Message string
}
