package usecase_regulation

type CreateRegulationInput struct {
	RegulationName string
	Abbreviation   string
}

type CreateRegulationOutput struct {
	RegulationID uint64
}
