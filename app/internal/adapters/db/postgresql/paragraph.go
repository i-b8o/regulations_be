package postgressql

import (
	"prod_serv/internal/domain/entity"
	client "prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient, logger *logging.Logger) paragraphStorage {
	return paragraphStorage{client: client}
}

// func (cs *paragraphStorage) GetOne(paragraphID string) *entity.Paragraph {
// 	return nil
// }

// func (cs *paragraphStorage) GetAll(chapterID string) []*entity.Paragraph {
// 	return nil
// }

func (cs *paragraphStorage) CreateAll(paragraphs []entity.Paragraph) error {
	return nil
}
