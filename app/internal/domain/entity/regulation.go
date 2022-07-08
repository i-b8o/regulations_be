package entity

import "time"

type Regulation struct {
	Id           uint64     `json:"id,omitempty"`
	Name         string     `json:"name"`
	Abbreviation string     `json:"abbreviation"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Chapters     []*Chapter `json:"chapters"`
}
