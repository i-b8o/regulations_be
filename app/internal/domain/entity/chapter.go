package entity

type Chapter struct {
	ID           uint64
	Name         string
	Num          string
	RegulationID uint64
	Paragraphs   []Paragraph
}
