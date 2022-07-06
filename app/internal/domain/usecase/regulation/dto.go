package usecase_regulation

type CreateRegulationInput struct {
	RegulationName string
}

type CreateRegulationOutput struct {
	RegulationID uint64
}
