package dto

import "fmt"

type CreateChapterRequest struct {
	RegulationID string `json:"regulation_id,omitempty"`
	ChapterName  string `json:"chapter_name"`
	ChapterNum   string `json:"chapter_num,omitempty"`
}

func (dto *CreateChapterRequest) Validate() error {
	if dto.ChapterName == "" {
		return fmt.Errorf("missing chapter name")
	}
	if dto.RegulationID == "" {
		return fmt.Errorf("missing regulation id")
	}

	return nil
}

type CreateChapterResponse struct {
	ChapterID string `json:"chapter_id"`
}
