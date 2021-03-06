package postgressql

import (
	"context"
	"errors"
	"fmt"
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

func (cs *chapterStorage) GetAllById(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Chapter) {
	const sql = `SELECT id,name,num,order_num FROM "chapters" WHERE r_id = $1 ORDER BY order_num`
	var resp entity.Response
	var chapters []*entity.Chapter

	rows, err := cs.client.Query(ctx, sql, regulationID)
	if err != nil {
		resp.Errors = append(resp.Errors, "Chapter GetAllByID Query "+err.Error())
		return resp, nil
	}

	defer rows.Close()

	for rows.Next() {
		chapter := &entity.Chapter{}
		if err = rows.Scan(
			&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum,
		); err != nil {
			resp.Errors = append(resp.Errors, "Chapter GetAllByID Next "+err.Error())
			return resp, nil
		}

		chapters = append(chapters, chapter)
	}

	return resp, chapters

}

func (cs *chapterStorage) Create(ctx context.Context, chapter entity.Chapter) entity.Response {
	sql := `INSERT INTO chapters ("name", "num", "order_num","r_id") VALUES ($1,$2,$3,$4) RETURNING "id"`
	fmt.Printf("order num %d", chapter.OrderNum)
	row := cs.client.QueryRow(ctx, sql, chapter.Name, chapter.Num, chapter.OrderNum, chapter.RegulationID)
	fmt.Println("aaaaaaaaaaaaaa", chapter.Name, chapter.Num, chapter.OrderNum, chapter.RegulationID)
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

func (cs *chapterStorage) GetOrderNum(ctx context.Context, id uint64) (orderNum uint64, err error) {
	const sql = `SELECT order_num FROM "chapters" WHERE id = $1 LIMIT 1`
	row := cs.client.QueryRow(ctx, sql, id)

	switch err := row.Scan(&orderNum); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		return 0, err
	case err != nil:
		log.Printf("cannot get regulation from database: %v\n", err)
		return 0, err
	}

	return orderNum, nil
}
