package models

type Paragraph struct {
	Id         int              `json:"id"`
	Text       []string         `json:"text"`
	EditedText []string         `json:"editedText,omitempty"`
	Tables     []ParagraphTable `json:"tables,omitempty"`
}
