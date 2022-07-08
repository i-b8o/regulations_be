package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ChapterStorage interface {
	Create(ctx context.Context, params entity.Chapter) entity.Response
	GetAllById(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Chapter)
}

type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s chapterService) GetAllById(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Chapter) {
	return s.storage.GetAllById(ctx, regulationID)
}

func (s chapterService) Create(ctx context.Context, chapter entity.Chapter) entity.Response {
	return s.storage.Create(ctx, chapter)

}
