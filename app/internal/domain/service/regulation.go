package service

import (
	"context"
	adapters_dto "prod_serv/internal/adapters/db/dto"
	"prod_serv/internal/domain/entity"
)

type RegulationStorage interface {
	Create(ctx context.Context, params adapters_dto.CreateRegulationParams) (uint64, error)
	// GetOne(id string) *entity.Regulation
	// GetAll() []*entity.RegulationNamesAndIDsView
}

type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) Create(ctx context.Context, regulation entity.Regulation) (uint64, error) {
	// MAPPING entity.Regulation --> adapters_dto.CreateRegulationParams
	adapterDTO := adapters_dto.CreateRegulationParams{
		RegulationName: regulation.RegulationName,
	}
	return s.storage.Create(ctx, adapterDTO)
}

// func (s *regulationService) GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return s.storage.GetAll()
// }

// func (s *regulationService) GetByID(ctx context.Context, id string) *entity.Regulation {
// 	return s.storage.GetOne(id)
// }
