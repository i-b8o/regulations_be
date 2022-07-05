package regulation_usecase

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationService interface {
	GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView
	GetByID(ctx context.Context, id string) *entity.Regulation
	CreateRegulation(ctx context.Context, reg CreateRegulationDTO) error
}

type regulationUsecase struct {
	regulationService RegulationService
}

func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
}

func (u regulationUsecase) CreateRegulation(ctx context.Context, reg CreateRegulationDTO) error {
	return u.regulationService.CreateRegulation(ctx, reg)
}

func (u regulationUsecase) GetRegulationByID(ctx context.Context, id string) *entity.Regulation {
	return u.regulationService.GetByID(ctx, id)
}
