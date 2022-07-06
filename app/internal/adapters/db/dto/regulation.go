package dto

import "time"

// CreateRegulationParams used by CreateRegulation.
type CreateRegulationParams struct {
	RegulationName string
	CreatedAt      time.Time
}
