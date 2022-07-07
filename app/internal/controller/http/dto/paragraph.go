package dto

import "fmt"

type CreateParagraphsRequest struct {
	Paragraphs []Paragraph `json:"paragraphs,omitempty"`
}

type Paragraph struct {
	ParagraphID       uint64 `json:"paragraph_id"`
	ParagraphOrderNum uint64 `json:"paragraph_order_num"`
	ParagraphClass    string `json:"paragraph_class,omitempty"`
	ParagraphText     string `json:"paragraph_text"`
	ChapterID         uint64 `json:"chapter_id"`
}

func (dto *Paragraph) Validate() error {

	if dto.ParagraphText == "" {
		return fmt.Errorf("missing paragraph text")
	}

	if dto.ParagraphClass == "" {
		dto.ParagraphClass = "-"
	}

	return nil
}

type CreateParagraphsResponse struct {
	Message string `json:"mesage,omitempty"`
}
