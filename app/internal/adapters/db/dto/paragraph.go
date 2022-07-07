package dto

import "prod_serv/internal/domain/entity"

// CreateParagraphsInput used by CreateParagraphs.
type CreateParagraphsInput struct {
	Paragraphs []entity.Paragraph
}

// CreateParagraphsResonse returned by CreateParagraphs.
type CreateParagraphsOutput struct {
	Message string
}