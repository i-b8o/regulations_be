package service

import (
	"prod_serv/internal/domain/entity"
)

type ChapterStorage interface {
	GetAll(regulationID string) []*entity.Chapter
	Create(chapter *entity.Chapter) error
}

type chapterService struct {
	storage ChapterStorage
}

func NewchapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s *chapterService) GetAllChaptersForRegulationID(regulationID string) []*entity.Chapter {
	return s.storage.GetAll(regulationID)
}
