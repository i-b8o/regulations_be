package service

import (
	"context"
	adapters_dto "prod_serv/internal/adapters/db/dto"
	"prod_serv/internal/domain/entity"
)

type ChapterStorage interface {
	Create(ctx context.Context, params adapters_dto.CreateChapterInput) (adapters_dto.CreateChapterOutput, error)
}

type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s chapterService) Create(ctx context.Context, chapter entity.Chapter) (entity.Chapter, error) {
	// MAPPING entity.Chapter --> adapters_dto.CreateChapterParams
	adapterDTO := adapters_dto.CreateChapterInput{
		ChapterName:  chapter.Name,
		ChapterNum:   chapter.Num,
		RegulationID: chapter.RegulationID,
	}
	adapterOut, err := s.storage.Create(ctx, adapterDTO)
	return entity.Chapter{ID: adapterOut.ChapterID}, err
}
