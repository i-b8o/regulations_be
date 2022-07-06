package entity

import "time"

type Regulation struct {
	RegulationId   int
	RegulationName string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

// type RegulationNamesAndIDsView struct {
// 	RegulationId   int
// 	RegulationName string
// }
