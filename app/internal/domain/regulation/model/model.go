package model

type Regulation struct {
	RegulationId   string `json:"regulation_id"`
	RegulationName string `json:"regulation_name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
