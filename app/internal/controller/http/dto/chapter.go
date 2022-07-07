package dto

import "fmt"

type CreateChapterRequest struct {
	RegulationID uint64 `json:"regulation_id"`
	ChapterName  string `json:"chapter_name"`
	ChapterNum   string `json:"chapter_num,omitempty"`
}

func (dto *CreateChapterRequest) Validate() error {
	if dto.ChapterName == "" {
		return fmt.Errorf("missing chapter name")
	}

	return nil
}

type CreateChapterResponse struct {
	ChapterID string `json:"chapter_id"`
}
