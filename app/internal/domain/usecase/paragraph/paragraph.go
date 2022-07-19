package usecase_paragraph

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ParagraphService interface {
	CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response
}

type LinkService interface {
	Create(ctx context.Context, params entity.Link) entity.Response
}

type paragraphUsecase struct {
	paragraphService ParagraphService
	linkService      LinkService
}

func NewParagraphUsecase(paragraphService ParagraphService, linkService LinkService) *paragraphUsecase {
	return &paragraphUsecase{paragraphService: paragraphService, linkService: linkService}
}

func (u paragraphUsecase) CreateParagraphs(ctx context.Context, paragraphs []entity.Paragraph) entity.Response {
	for _, p := range paragraphs {
		if p.ID > 0 {
			u.linkService.Create(ctx, entity.Link{ID: p.ID, ParagraphNum: p.Num, ChapterID: p.ChapterID})
		}
	}
	return u.paragraphService.CreateAll(ctx, paragraphs)

}
