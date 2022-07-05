package postgressql

import (
	"prod_serv/internal/domain/entity"
	"prod_serv/pkg/logging"

	sq "github.com/Masterminds/squirrel"
)

type regulationStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewRegulationStorage(client PostgreSQLClient, logger *logging.Logger) regulationStorage {
	return regulationStorage{queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), client: client}
}

func (rs *regulationStorage) Create(regulation *entity.Regulation) error {
	return nil
}

func (rs *regulationStorage) GetOne(id string) *entity.Regulation {
	return nil
}

func (rs *regulationStorage) GetAll() []*entity.RegulationNamesAndIDsView {
	return nil
}
