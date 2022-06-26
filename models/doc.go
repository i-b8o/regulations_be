package models

type Doc struct {
	Abbreviation string    `json:"abbreviation"`
	Name         string    `json:"name"`
	Chapters     []Chapter `json:"chapters"`
}
