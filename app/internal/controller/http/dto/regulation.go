package dto

import (
	"fmt"
	"prod_serv/internal/domain/entity"
)

type CreateRegulationRequestDTO struct {
	RegulationName string `json:"regulation_name"`
	Abbreviation   string `json:"abbreviation"`
}

func (dto *CreateRegulationRequestDTO) Validate() error {
	if dto.RegulationName == "" {
		return fmt.Errorf("missing regulation name")
	}

	return nil
}

type GetFullRegulationRequestDTO struct {
	RegulationID uint64 `json:"regulation_id"`
}

type GetFullRegulationResponseDTO struct {
	Response       entity.Response `json:"response"`
	RegulationID   uint64          `json:"regulation_id"`
	RegulationName string          `json:"regulation_name"`
	Abbreviation   string          `json:"abbreviation"`
}
