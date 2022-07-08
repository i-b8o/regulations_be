package service

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationStorage interface {
	Create(ctx context.Context, params entity.Regulation) entity.Response
	GetOne(ctx context.Context, regulation entity.Regulation) (entity.Response, entity.Regulation)
	// GetOne(id string) *entity.Regulation
	// GetAll() []*entity.RegulationNamesAndIDsView
}

type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) Create(ctx context.Context, regulation entity.Regulation) entity.Response {
	return s.storage.Create(ctx, regulation)
}

func (s *regulationService) GetOne(ctx context.Context, regulation entity.Regulation) (entity.Response, entity.Regulation) {
	return s.storage.GetOne(ctx, regulation)
}

// func (s *regulationService) GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return s.storage.GetAll()
// }

// func (s *regulationService) GetByID(ctx context.Context, id string) *entity.Regulation {
// 	return s.storage.GetOne(id)
// }
