package models

type Paragraph struct {
	Chapter    string           `json:"chapter"`
	Text       []string         `json:"text"`
	EditedText []string         `json:"editedText"`
	Tables     []ParagraphTable `json:"tables"`
}
