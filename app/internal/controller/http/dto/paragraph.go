package dto

type Paragraph struct {
	ID                string `json:"id,omitempty"`
	Href              string `json:"href,omitempty"`
	ParagraphOrderNum string `json:"paragraph_order_num"`
	ParagraphClass    string `json:"p_class,omitempty"`
	ParagraphText     string `json:"paragraph_text"`
}
