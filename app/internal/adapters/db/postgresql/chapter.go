package postgressql

import (
	"context"
	"errors"
	"prod_serv/internal/adapters/db/dto"

	client "prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient, logger *logging.Logger) *chapterStorage {
	return &chapterStorage{client: client}
}

func (cs *chapterStorage) Create(ctx context.Context, params dto.CreateChapterInput) (dto.CreateChapterOutput, error) {
	sql := `INSERT INTO chapters ("name", "num", "r_id") VALUES ($1,$2,$3) RETURNING "id"`

	row := cs.client.QueryRow(ctx, sql, params.ChapterName, params.ChapterNum, params.RegulationID)
	var chapterID uint64
	resp := dto.CreateChapterOutput{}
	switch err := row.Scan(&chapterID); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		return resp, err
	case err != nil:
		if sqlErr := cs.chapterPgError(err); sqlErr != nil {
			return resp, sqlErr
		}
		return resp, errors.New("cannot create chapter on database")
	}
	resp.ChapterID = chapterID
	return resp, nil
}

func (cs *chapterStorage) chapterPgError(err error) error {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return nil
	}
	if pgErr.Code == pgerrcode.UniqueViolation {
		return errors.New("chapter already exists")
	}
	if pgErr.Code == pgerrcode.CheckViolation {
		switch pgErr.ConstraintName {
		case "name_check":
			return errors.New("invalid chapter name")
		default:
			return errors.New("invalid ")
		}
	}
	return nil
}

// func (cs *chapterStorage) GetAll(regulationID string) []*entity.Chapter {
// 	return nil
// }

// func (cs *chapterStorage) GetOne(id string) *entity.Chapter {
// 	return nil
// }
