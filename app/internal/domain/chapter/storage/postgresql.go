package storage

import (
	"context"

	"prod_serv/internal/domain/chapter/model"
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
	table  = "chapter"
)

func (s *ChapterStorage) All(ctx context.Context) ([]model.Chapter, error) {
	s.queryBuilder.Select("id").
		Column("name").
		Column("num").
		Column("created_at").
		Column("updated_at")
	return nil, nil
}
