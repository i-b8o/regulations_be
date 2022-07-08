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

type GetFullRegulationJSONResponseDTO struct {
	Response   entity.Response   `json:"response"`
	Regulation entity.Regulation `json:"regulation"`
}

type GetFullRegulationDartResponseDTO struct {
	Response   entity.Response `json:"response"`
	Regulation string          `json:"regulation"`
}
