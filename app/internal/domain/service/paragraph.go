package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ParagraphStorage interface {
	CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response
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
