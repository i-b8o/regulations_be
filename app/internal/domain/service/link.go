package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type LinkStorage interface {
	Create(ctx context.Context, params entity.Link) entity.Response
	GetAllByChapterID(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Link)
	GetAll(ctx context.Context) (entity.Response, []*entity.Link)
}

type linkService struct {
	storage LinkStorage
}

func NewLinkService(storage LinkStorage) *linkService {
	return &linkService{storage: storage}
}

func (s linkService) GetAllByChapterID(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Link) {
	return s.storage.GetAllByChapterID(ctx, regulationID)
}

func (s linkService) GetAll(ctx context.Context) (entity.Response, []*entity.Link) {
	return s.storage.GetAll(ctx)
}

func (s linkService) Create(ctx context.Context, chapter entity.Link) entity.Response {
	return s.storage.Create(ctx, chapter)

}
