package usecase_chapter

import (
	"context"
	"prod_serv/internal/domain/entity"
	"strconv"
)

type ChapterService interface {
	Create(ctx context.Context, chapter entity.Chapter) entity.Response
}

type LinkService interface {
	Create(ctx context.Context, params entity.Link) entity.Response
}

type chapterUsecase struct {
	chapterService ChapterService
	linkService    LinkService
}

func NewChapterUsecase(chapterService ChapterService, linkService LinkService) *chapterUsecase {
	return &chapterUsecase{chapterService: chapterService, linkService: linkService}
}

func (u chapterUsecase) CreateChapter(ctx context.Context, chapter entity.Chapter) entity.Response {
	resp := u.chapterService.Create(ctx, chapter)

	if chapter.ID > 0 {

		ch_id, err := strconv.ParseUint(resp.ID, 10, 64)

		if err != nil {

			resp.Errors = append(resp.Errors, err.Error())
			return resp
		}

		r := u.linkService.Create(ctx, entity.Link{ID: chapter.ID, ParagraphNum: 0, RID: chapter.RegulationID, ChapterID: ch_id})
		resp.Errors = append(resp.Errors, r.Errors...)

	}
	return resp
}
