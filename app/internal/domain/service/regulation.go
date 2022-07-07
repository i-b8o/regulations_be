package service

import (
	"context"
	adapters_dto "prod_serv/internal/adapters/db/dto"
	"prod_serv/internal/domain/entity"
)

type RegulationStorage interface {
	Create(ctx context.Context, params adapters_dto.CreateRegulationInput) (adapters_dto.CreateRegulationOutput, error)
	// GetOne(id string) *entity.Regulation
	// GetAll() []*entity.RegulationNamesAndIDsView
}

type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) Create(ctx context.Context, regulation entity.Regulation) (entity.Regulation, error) {
	// MAPPING entity.Regulation --> adapters_dto.CreateRegulationParams
	adapterDTO := adapters_dto.CreateRegulationInput{
		RegulationName: regulation.Name,
		Abbreviation:   regulation.Abbreviation,
	}
	out, err := s.storage.Create(ctx, adapterDTO)
	reg := entity.Regulation{Id: out.RegulationID}
	return reg, err
}

// func (s *regulationService) GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return s.storage.GetAll()
// }

// func (s *regulationService) GetByID(ctx context.Context, id string) *entity.Regulation {
// 	return s.storage.GetOne(id)
// }
