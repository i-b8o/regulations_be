package model

import "time"

type Regulation struct {
	RegulationId   int        `json:"regulation_id"`
	RegulationName string     `json:"regulation_name"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
