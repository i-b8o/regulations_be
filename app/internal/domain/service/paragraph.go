package service

import (
	"prod_serv/internal/domain/entity"
)

type ParagraphStorage interface {
	GetParagraphByID(paragraphID string) *entity.Paragraph
	GetAllParagraphsByChapterID(chapterID string) []entity.Paragraph
}

type paragraphService struct {
	storage ParagraphStorage
}

func NewparagraphService(storage ParagraphStorage) *paragraphService {
	return &paragraphService{storage: storage}
}

func (s *paragraphService) GetParagraphByID(paragraphID string) *entity.Paragraph {
	return s.storage.GetParagraphByID(paragraphID)
}

func (s *paragraphService) GetAllParagraphsByChapterID(chapterID string) []entity.Paragraph {
	return s.storage.GetAllParagraphsByChapterID(chapterID)
}
