package dto

import "fmt"

type ParagraphDTO struct {
	ID                string `json:"id,omitempty"`
	Href              string `json:"href,omitempty"`
	ParagraphOrderNum string `json:"paragraph_order_num"`
	ParagraphClass    string `json:"p_class,omitempty"`
	ParagraphText     string `json:"paragraph_text"`
}

func (dto *ParagraphDTO) Validate() error {
	if dto.ParagraphOrderNum == "" {
		return fmt.Errorf("missing paragraph num")
	}

	return nil
}
