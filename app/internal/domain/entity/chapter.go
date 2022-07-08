package entity

type Chapter struct {
	ID           uint64      `json:"id,omitempty"`
	Name         string      `json:"name"`
	Num          string      `json:"num,omitempty"`
	RegulationID uint64      `json:"regulation_id,omitempty"`
	Paragraphs   []Paragraph `json:"paragraphs,omitempty"`
}
