package entity

import "time"

type Regulation struct {
	RegulationId   int        `json:"regulation_id"`
	RegulationName string     `json:"regulation_name"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	Chapters       *[]Chapter `json:"chapters"`
}

type RegulationNamesAndIDsView struct {
	RegulationId   int    `json:"regulation_id"`
	RegulationName string `json:"regulation_name"`
}
