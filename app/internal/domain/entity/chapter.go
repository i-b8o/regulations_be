package entity

type Chapter struct {
	ChapterID   string       `json:"chapter_id,omitempty"`
	ChapterName string       `json:"chapter_name"`
	ChapterNum  string       `json:"chapter_num,omitempty"`
	Paragraphs  *[]Paragraph `json:"paragraphs"`
}
