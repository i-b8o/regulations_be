package dto

import "fmt"

type CreateChapterRequest struct {
	ChapterID    uint64 `json:"chapter_id,omitempty"`
	RegulationID uint64 `json:"regulation_id"`
	ChapterName  string `json:"chapter_name"`
	ChapterNum   string `json:"chapter_num,omitempty"`
	OrderNum     uint64 `json:"order_num"`
}

func (dto *CreateChapterRequest) Validate() (string, error) {
	if dto.ChapterName == "" {
		return "", fmt.Errorf("missing chapter name")
	}
	if dto.OrderNum < 1 {
		return "missing order num", nil
	}
	if dto.ChapterNum == "" {
		return "missing chapter num", nil
	}

	return "", nil
}
