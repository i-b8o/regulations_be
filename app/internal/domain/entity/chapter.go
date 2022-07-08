package entity

type Chapter struct {
	ID           uint64      `json:"chapter_id,omitempty"`
	Name         string      `json:"chapter_name"`
	Num          string      `json:"chapter_num,omitempty"`
	RegulationID uint64      `json:"regulation_num,omitempty"`
	Paragraphs   []Paragraph `json:"paragraphs,omitempty"`
}
