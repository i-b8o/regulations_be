package postgressql

import (
	"prod_serv/internal/domain/entity"
	"prod_serv/pkg/logging"

	client "prod_serv/pkg/client/postgresql"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient, logger *logging.Logger) chapterStorage {
	return chapterStorage{client: client}
}

func (cs *chapterStorage) GetAll(regulationID string) []*entity.Chapter {
	return nil
}

func (cs *chapterStorage) Create(chapter *entity.Chapter) error {
	return nil
}

// func (cs *chapterStorage) GetOne(id string) *entity.Chapter {
// 	return nil
// }
