package dto

import "fmt"

type CreateParagraphsRequest struct {
	Paragraphs []Paragraph `json:"paragraphs,omitempty"`
}

type Paragraph struct {
	ParagraphID       uint64 `json:"paragraph_id"`
	ParagraphOrderNum uint64 `json:"paragraph_order_num"`
	IsHTML            bool   `json:"is_html"`
	IsTable           bool   `json:"is_table"`
	IsNFT             bool   `json:"is_nft"`
	ParagraphClass    string `json:"paragraph_class,omitempty"`
	ParagraphText     string `json:"paragraph_text"`
	ChapterID         uint64 `json:"chapter_id"`
}

func (dto *Paragraph) Validate() (string, error) {

	if dto.ParagraphText == "" {
		return "", fmt.Errorf("missing paragraph text")
	}

	if dto.ParagraphClass == "" {
		dto.ParagraphClass = "-"
	}

	if dto.IsHTML {
		return fmt.Sprintf("Chapter ID:%d: HTML in %d", dto.ChapterID, dto.ParagraphOrderNum), nil
	}

	return "", nil
}
