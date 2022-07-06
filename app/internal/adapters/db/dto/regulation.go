package dto

import "time"

// CreateRegulationInput used by CreateRegulation.
type CreateRegulationInput struct {
	RegulationName string
	CreatedAt      time.Time
}

// CreateRegulationResonse returned by CreateRegulation.
type CreateRegulationOutput struct {
	RegulationID uint64
}
