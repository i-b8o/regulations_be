package service

import (
	"context"
	adapters_dto "prod_serv/internal/adapters/db/dto"
	"prod_serv/internal/domain/entity"
)

type ParagraphStorage interface {
	// GetOne(paragraphID string) *entity.Paragraph
	// GetAll(chapterID string) []*entity.Paragraph
	// CreateAll(paragraphs []entity.Paragraph) error
	CreateAll(ctx context.Context, in adapters_dto.CreateParagraphsInput) adapters_dto.CreateParagraphsOutput
}

type paragraphService struct {
	storage ParagraphStorage
}

func NewParagraphService(storage ParagraphStorage) *paragraphService {
	return &paragraphService{storage: storage}
}

// func (s *paragraphService) GetParagraphByID(paragraphID string) *entity.Paragraph {
// 	return s.storage.GetOne(paragraphID)
// }

// func (s *paragraphService) GetAllParagraphsByChapterID(chapterID string) []*entity.Paragraph {
// 	return s.storage.GetAll(chapterID)
// }

func (s *paragraphService) CreateAll(ctx context.Context, odt entity.ParagraphsInput) entity.ParagraphsOutput {
	in := adapters_dto.CreateParagraphsInput{
		Paragraphs: odt.Paragraphs,
	}
	out := s.storage.CreateAll(ctx, in)
	return entity.ParagraphsOutput{Message: out.Message}
}
