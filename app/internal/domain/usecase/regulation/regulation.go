package regulation

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationService interface {
	GetNamesAndIDsOfAllRegulations(ctx context.Context) []entity.RegulationNamesAndIDsView
	GetByID(ctx context.Context, id string) *entity.Regulation
}

type ChapterService interface {
	GetByID(ctx context.Context, regulationID int) []entity.Chapter
}

type ParagraphService interface {
	GetByID(ctx context.Context, chapterID string) []entity.Paragraph
}

type regulationUsecase struct {
	regulationService RegulationService
	chapterService    ChapterService
	paragraphService  ParagraphService
}

func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []entity.RegulationNamesAndIDsView {
	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
}

func (u regulationUsecase) GetRegulationByID(ctx context.Context, id string) entity.Regulation {
	regulation := u.regulationService.GetByID(ctx, id)
	chapters := u.chapterService.GetByID(ctx, regulation.RegulationId)
	for _, chapter := range chapters {
		paragraphs := u.paragraphService.GetByID(ctx, chapter.ChapterID)
		chapter.Paragraphs = paragraphs
		regulation.Chapters = append(regulation.Chapters, chapter)
	}
}
