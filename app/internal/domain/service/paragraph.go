package service

import (
	"prod_serv/internal/domain/entity"
)

type ParagraphStorage interface {
	GetOne(paragraphID string) *entity.Paragraph
	GetAll(chapterID string) []*entity.Paragraph
	CreateAll(paragraphs []entity.Paragraph) error
}

type paragraphService struct {
	storage ParagraphStorage
}

func NewparagraphService(storage ParagraphStorage) *paragraphService {
	return &paragraphService{storage: storage}
}

func (s *paragraphService) GetParagraphByID(paragraphID string) *entity.Paragraph {
	return s.storage.GetOne(paragraphID)
}

func (s *paragraphService) GetAllParagraphsByChapterID(chapterID string) []*entity.Paragraph {
	return s.storage.GetAll(chapterID)
}

func (s *paragraphService) CreateAllParagraphs(paragraphs []entity.Paragraph) error {
	return s.storage.CreateAll(paragraphs)
}
