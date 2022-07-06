package dto

import "fmt"

type CreateChapterDTO struct {
	ChapterID   string `json:"chapter_id,omitempty"`
	ChapterName string `json:"chapter_name"`
	ChapterNum  string `json:"chapter_num,omitempty"`
}

func (dto *CreateChapterDTO) Validate() error {
	if dto.ChapterName == "" {
		return fmt.Errorf("missing chapter name")
	}

	return nil
}
