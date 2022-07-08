package usecase_regulation

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationService interface {
	// GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetByID(ctx context.Context, id string) *entity.Regulation
	GetOne(ctx context.Context, regulation entity.Regulation) (entity.Response, entity.Regulation)
	Create(ctx context.Context, regulation entity.Regulation) entity.Response
}

type regulationUsecase struct {
	regulationService RegulationService
}

func NewRegulationUsecase(regulationService RegulationService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService}
}

func (u regulationUsecase) CreateRegulation(ctx context.Context, regulation entity.Regulation) entity.Response {
	return u.regulationService.Create(ctx, regulation)
}

func (u regulationUsecase) GetFullRegulationByID(ctx context.Context, regulation entity.Regulation) (entity.Response, entity.Regulation) {
	return u.regulationService.GetOne(ctx, regulation)
}

// func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
// }
