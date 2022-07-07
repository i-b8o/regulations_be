package entity

import "time"

type Regulation struct {
	Id           uint64
	Name         string
	Abbreviation string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

// type RegulationNamesAndIDsView struct {
// 	RegulationId   int
// 	RegulationName string
// }
