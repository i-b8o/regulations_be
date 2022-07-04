package model

type Chapter struct {
	ChapterID   string `json:"chapter_id"`
	ChapterName string `json:"chapter_name"`
	ChapterNum  string `json:"chapter_num"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
