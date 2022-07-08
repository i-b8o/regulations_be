package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ChapterStorage interface {
	Create(ctx context.Context, params entity.Chapter) entity.Response
}

type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s chapterService) Create(ctx context.Context, chapter entity.Chapter) entity.Response {
	return s.storage.Create(ctx, chapter)

}
