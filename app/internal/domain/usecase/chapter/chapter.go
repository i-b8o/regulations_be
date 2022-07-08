package usecase_chapter

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ChapterService interface {
	Create(ctx context.Context, chapter entity.Chapter) entity.Response
}

type chapterUsecase struct {
	regulationService ChapterService
}

func NewChapterUsecase(regulationService ChapterService) *chapterUsecase {
	return &chapterUsecase{regulationService: regulationService}
}

func (u chapterUsecase) CreateChapter(ctx context.Context, chapter entity.Chapter) entity.Response {
	return u.regulationService.Create(ctx, chapter)
}
