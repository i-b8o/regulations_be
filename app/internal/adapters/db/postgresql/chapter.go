package postgressql

import (
	"context"
	"errors"
	"log"
	"strconv"

	"prod_serv/internal/domain/entity"
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

func (cs *chapterStorage) GetOne(ctx context.Context, regulation entity.Regulation) (entity.Response, []entity.Chapter) {
	const sql = `SELECT id,name,num FROM "chapters" WHERE r_id = $1 LIMIT 1`
	row := cs.client.QueryRow(ctx, sql, regulation.Id)

	var resp entity.Response
	var chapter entity.Chapter
	switch err := row.Scan(&chapter.ID, &chapter.Name, &chapter.Num); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		resp.Errors = append(resp.Errors, err.Error())
		return resp, chapter
	case err != nil:
		resp.Errors = append(resp.Errors, err.Error())
		log.Printf("cannot get chapter from database: %v\n", err)
		return resp, chapter
	}

	return resp, chapter
}

func (cs *chapterStorage) Create(ctx context.Context, chapter entity.Chapter) entity.Response {
	sql := `INSERT INTO chapters ("name", "num", "r_id") VALUES ($1,$2,$3) RETURNING "id"`

	row := cs.client.QueryRow(ctx, sql, chapter.Name, chapter.Num, chapter.RegulationID)
	var chapterID uint64
	resp := entity.Response{}
	switch err := row.Scan(&chapterID); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		resp.Errors = append(resp.Errors, err.Error())
		return resp
	case err != nil:
		if sqlErr := cs.chapterPgError(err); sqlErr != nil {
			resp.Errors = append(resp.Errors, sqlErr.Error())
			return resp
		}
		resp.Errors = append(resp.Errors, err.Error())
		return resp
	}
	resp.ID = strconv.FormatUint(chapterID, 10)
	return resp
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
