package usecase_regulation

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationService interface {
	GetOne(ctx context.Context, regulationId uint64) (entity.Response, entity.Regulation)
	Create(ctx context.Context, regulation entity.Regulation) entity.Response
}
type ChapterService interface {
	GetAllById(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Chapter)
}
type ParagraphService interface {
	GetAllById(ctx context.Context, chapterID uint64) (entity.Response, []entity.Paragraph)
}
type regulationUsecase struct {
	regulationService RegulationService
	chapterService    ChapterService
	paragraphService  ParagraphService
}

func NewRegulationUsecase(regulationService RegulationService, chapterService ChapterService, paragraphService ParagraphService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService, chapterService: chapterService, paragraphService: paragraphService}
}

func (u regulationUsecase) CreateRegulation(ctx context.Context, regulation entity.Regulation) entity.Response {
	return u.regulationService.Create(ctx, regulation)
}

func (u regulationUsecase) GetFullRegulationByID(ctx context.Context, regulationID uint64) (entity.Response, entity.Regulation) {
	resp, regulation := u.regulationService.GetOne(ctx, regulationID)
	respErrors := resp.Errors
	resp, chapters := u.chapterService.GetAllById(ctx, regulationID)
	resp.Errors = append(resp.Errors, respErrors...)

	for _, chapter := range chapters {
		response, paragraphs := u.paragraphService.GetAllById(ctx, chapter.ID)
		if len(response.Errors) > 0 {
			resp.Errors = append(resp.Errors, response.Errors...)
		}

		chapter.Paragraphs = paragraphs
	}

	// fmt.Println(len(chapters[0].Paragraphs))

	regulation.Chapters = chapters

	return resp, regulation
}

// func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
// }
