package usecase_chapter

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ChapterService interface {
	// GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetByID(ctx context.Context, id string) *entity.Regulation

	Create(ctx context.Context, chapter entity.Chapter) (entity.Chapter, error)
}

type chapterUsecase struct {
	regulationService ChapterService
}

func NewChapterUsecase(regulationService ChapterService) *chapterUsecase {
	return &chapterUsecase{regulationService: regulationService}
}

// func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
// }

func (u chapterUsecase) CreateChapter(ctx context.Context, dto CreateChapterInput) (CreateChapterOutput, error) {
	// MAPPING dto.CreateChapterDTO --> entity.Chapter
	entityDTO := entity.Chapter{
		Name:         dto.Name,
		Num:          dto.Num,
		RegulationID: dto.RegulationID,
	}

	chapter, err := u.regulationService.Create(ctx, entityDTO)
	return CreateChapterOutput{ChapterID: chapter.ID}, err
}

// func (u regulationUsecase) GetRegulationByID(ctx context.Context, id string) *entity.Regulation {
// 	return u.regulationService.GetByID(ctx, id)
// }
