package usecase_paragraph

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ParagraphService interface {
	CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response
}

type paragraphUsecase struct {
	paragraphService ParagraphService
}

func NewParagraphUsecase(paragraphService ParagraphService) *paragraphUsecase {
	return &paragraphUsecase{paragraphService: paragraphService}
}

func (u paragraphUsecase) CreateParagraphs(ctx context.Context, paragraphs []entity.Paragraph) entity.Response {
	return u.paragraphService.CreateAll(ctx, paragraphs)

}
