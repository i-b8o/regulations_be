package models

type Chapter struct {
	Name       string      `json:"name"`
	Num        string      `json:"num"`
	Paragraphs []Paragraph `json:"paragraphs"`
}
