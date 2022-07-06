package dto

import "fmt"

type CreateRegulationRequestDTO struct {
	RegulationName string `json:"regulation_name"`
}

func (dto *CreateRegulationRequestDTO) Validate() error {
	if dto.RegulationName == "" {
		return fmt.Errorf("missing regulation name")
	}

	return nil
}

type CreateRegulationResponseDTO struct {
	RegulationID string `json:"regulation_id"`
}
