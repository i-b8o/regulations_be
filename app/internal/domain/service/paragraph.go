package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ParagraphStorage interface {
	CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response
	GetAllById(ctx context.Context, chapterID uint64) (entity.Response, []entity.Paragraph)
}

type paragraphService struct {
	storage ParagraphStorage
}

func NewParagraphService(storage ParagraphStorage) *paragraphService {
	return &paragraphService{storage: storage}
}

func (s *paragraphService) CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response {
	return s.storage.CreateAll(ctx, paragraphs)
}

func (s *paragraphService) GetAllById(ctx context.Context, chapterID uint64) (entity.Response, []entity.Paragraph) {
	return s.storage.GetAllById(ctx, chapterID)
}
