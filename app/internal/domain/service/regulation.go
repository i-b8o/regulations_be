package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationStorage interface {
	Create(regulation *entity.Regulation) error
	GetNamesAndIDsOfAllRegulations() []*entity.RegulationNamesAndIDsView
	GetOne(id string) *entity.Regulation
}

type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) Create(ctx context.Context, regulation *entity.Regulation) error {
	return s.storage.Create(regulation)
}

func (s *regulationService) GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView {
	return s.storage.GetNamesAndIDsOfAllRegulations()
}

func (s *regulationService) GetByID(ctx context.Context, id string) *entity.Regulation {
	return s.storage.GetOne(id)
}
