package regulation_usecase

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationService interface {
	// GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetByID(ctx context.Context, id string) *entity.Regulation

	Create(ctx context.Context, regulation entity.Regulation) (uint64, error)
}

type regulationUsecase struct {
	regulationService RegulationService
}

func NewRegulationUsecase(regulationService RegulationService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService}
}

// func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
// }

func (u regulationUsecase) CreateRegulation(ctx context.Context, dto CreateRegulationDTO) (uint64, error) {
	// MAPPING dto.CreateRegulationDTO --> entity.CreateRegulationDTO
	entityDTO := entity.Regulation{
		RegulationName: dto.RegulationName,
	}
	return u.regulationService.Create(ctx, entityDTO)
}

// func (u regulationUsecase) GetRegulationByID(ctx context.Context, id string) *entity.Regulation {
// 	return u.regulationService.GetByID(ctx, id)
// }
