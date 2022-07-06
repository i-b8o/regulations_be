package usecase_regulation

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type RegulationService interface {
	// GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetByID(ctx context.Context, id string) *entity.Regulation

	Create(ctx context.Context, regulation entity.Regulation) (entity.Regulation, error)
}

type regulationUsecase struct {
	regulationService RegulationService
}

func NewRegulationUsecase(regulationService RegulationService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService}
}

func (u regulationUsecase) CreateRegulation(ctx context.Context, dto CreateRegulationInput) (CreateRegulationOutput, error) {
	// MAPPING dto.CreateRegulationDTO --> entity.Regulation
	regulation := entity.Regulation{
		Name: dto.RegulationName,
	}

	reg, err := u.regulationService.Create(ctx, regulation)
	return CreateRegulationOutput{RegulationID: reg.Id}, err
}

// func (u regulationUsecase) GetRegulationByID(ctx context.Context, id string) *entity.Regulation {
// 	return u.regulationService.GetByID(ctx, id)
// }

// func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
// }
