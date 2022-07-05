package postgressql

import (
	"prod_serv/internal/domain/entity"
	"prod_serv/pkg/logging"

	sq "github.com/Masterminds/squirrel"
)

type chapterStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewChapterStorage(client PostgreSQLClient, logger *logging.Logger) chapterStorage {
	return chapterStorage{queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), client: client}
}

func (cs *chapterStorage) GetAllChaptersByRegulationID(regulationID string) []*entity.Chapter {
	return nil
}
