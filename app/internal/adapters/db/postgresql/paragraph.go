package postgressql

import (
	"prod_serv/internal/domain/entity"
	"prod_serv/pkg/logging"

	sq "github.com/Masterminds/squirrel"
)

type paragraphStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewParagraphStorage(client PostgreSQLClient, logger *logging.Logger) paragraphStorage {
	return paragraphStorage{queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), client: client}
}

func (cs *paragraphStorage) GetOne(paragraphID string) *entity.Paragraph {
	return nil
}

func (cs *paragraphStorage) GetAll(chapterID string) []*entity.Paragraph {
	return nil
}

func (cs *paragraphStorage) CreateAll(paragraphs []entity.Paragraph) error {
	return nil
}
