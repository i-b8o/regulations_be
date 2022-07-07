package dto

import "fmt"

type CreateParagraphsRequest struct {
	Paragraphs []Paragraph `json:"paragraphs,omitempty"`
}

type Paragraph struct {
	paragraphID       string `json:"paragraph_id"`
	ParagraphOrderNum string `json:"paragraph_order_num"`
	ParagraphClass    string `json:"paragraph_class,omitempty"`
	ParagraphText     string `json:"paragraph_text"`
	ChapterID         string `json:"chapter_id"`
}

func (dto *Paragraph) Validate() error {
	if dto.ParagraphOrderNum == "" {
		return fmt.Errorf("missing paragraph num")
	}

	if dto.ParagraphText == "" {
		return fmt.Errorf("missing paragraph text")
	}

	if dto.ChapterID == "" {
		return fmt.Errorf("missing chapter id")
	}

	if dto.ParagraphClass == "" {
		dto.ParagraphClass = "-"
	}

	return nil
}

type CreateParagraphsResponse struct {
	Message string `json:"mesage,omitempty"`
}
