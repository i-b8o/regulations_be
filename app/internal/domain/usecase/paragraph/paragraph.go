package usecase_paragraph

import (
	"context"
	"prod_serv/internal/domain/entity"
)

type ParagraphService interface {
	// GetNamesAndIDsOfAllRegulations(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetByID(ctx context.Context, id string) *entity.Regulation

	CreateAll(ctx context.Context, paragraphs entity.ParagraphsInput) entity.ParagraphsOutput
}

type paragraphUsecase struct {
	paragraphService ParagraphService
}

func NewParagraphUsecase(paragraphService ParagraphService) *paragraphUsecase {
	return &paragraphUsecase{paragraphService: paragraphService}
}

func (u paragraphUsecase) CreateParagraphs(ctx context.Context, dto CreateParagraphsInput) CreateParagraphsOutput {
	// MAPPING dto.CreateParagraphsInput --> []entity.Paragraph
	serviceIn := entity.ParagraphsInput{
		Paragraphs: dto.Paragraphs,
	}

	serviceOut := u.paragraphService.CreateAll(ctx, serviceIn)
	return CreateParagraphsOutput{Message: serviceOut.Message}
}

// func (u paragraphUsecase) GetRegulationByID(ctx context.Context, id string) *entity.Regulation {
// 	return u.paragraphService.GetByID(ctx, id)
// }

// func (u paragraphUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.paragraphService.GetNamesAndIDsOfAllRegulations(ctx)
// }
