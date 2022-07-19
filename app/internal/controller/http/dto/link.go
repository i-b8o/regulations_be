package dto

import "prod_serv/internal/domain/entity"

type CreateLinkRequest struct {
	ID           uint64 `json:"id"`
	ChapterID    uint64 `json:"chapter_id"`
	ParagraphNum uint64 `json:"paragraph_num"`
}

func (dto *CreateLinkRequest) Validate() (string, error) {
	if dto.ID < 1 {
		return "missing link id", nil
	}
	if dto.ChapterID < 1 {
		return "missing chapter num", nil
	}

	if dto.ParagraphNum < 1 {
		return "missing paragraph num", nil
	}

	return "", nil
}

type GetAllLinksRequestDTO struct {
	RegulationID uint64 `json:"regulation_id"`
}

type GetAllLinksDartResponseDTO struct {
	Response entity.Link `json:"response"`
	AllLink  string      `json:"all_links"`
}
