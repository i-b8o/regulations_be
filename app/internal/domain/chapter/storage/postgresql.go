package storage

import (
	"prod_serv/pkg/logging"

	sq "github.com/Masterminds/squirrel"
)

type ChapterStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
	logger       *logging.Logger
}

func NewChapterStorage(client PostgreSQLClient, logger *logging.Logger) ChapterStorage {

	return ChapterStorage{queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), client: client, logger: logger}
}

const (
	scheme = "public"
	table  = "chapters"
)
