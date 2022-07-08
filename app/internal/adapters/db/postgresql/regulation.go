package postgressql

import (
	"context"
	"errors"
	"log"
	"prod_serv/internal/domain/entity"
	"strconv"
	"time"

	client "prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type regulationStorage struct {
	client client.PostgreSQLClient
}

func NewRegulationStorage(client client.PostgreSQLClient, logger *logging.Logger) *regulationStorage {
	return &regulationStorage{client: client}
}

func (rs *regulationStorage) GetOne(ctx context.Context, regulationID uint64) (entity.Response, entity.Regulation) {
	const sql = `SELECT name,abbreviation FROM "regulations" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, regulationID)

	var resp entity.Response
	var regulation entity.Regulation
	switch err := row.Scan(&regulation.Name, &regulation.Abbreviation); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		resp.Errors = append(resp.Errors, "Regulation GetOne Scan "+err.Error())
		return resp, regulation
	case err != nil:
		resp.Errors = append(resp.Errors, "Regulation GetOne Scan "+err.Error())
		log.Printf("cannot get regulation from database: %v\n", err)
		return resp, regulation
	}

	return resp, regulation
}

func (rs *regulationStorage) Create(ctx context.Context, regulation entity.Regulation) entity.Response {
	t := time.Now()

	const sql = `INSERT INTO regulations ("name", "abbreviation", "created_at") VALUES ($1, $2, $3) RETURNING "id"`

	row := rs.client.QueryRow(ctx, sql, regulation.Name, regulation.Abbreviation, t)
	var regulationID uint64
	resp := entity.Response{}
	switch err := row.Scan(&regulationID); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		resp.Errors = append(resp.Errors, err.Error())
		return resp
	case err != nil:
		if sqlErr := rs.createPgError(err); sqlErr != nil {
			resp.Errors = append(resp.Errors, sqlErr.Error())
			return resp
		}
		resp.Errors = append(resp.Errors, err.Error())
		log.Printf("cannot create regulation on database: %v\n", err)
		return resp
	}
	resp.ID = strconv.FormatUint(regulationID, 10)
	return resp
}

func (rs *regulationStorage) createPgError(err error) error {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return nil
	}
	if pgErr.Code == pgerrcode.UniqueViolation {
		return errors.New("regulation already exists")
	}
	if pgErr.Code == pgerrcode.CheckViolation {
		switch pgErr.ConstraintName {
		case "name_check":
			return errors.New("invalid regulation name")
		default:
			return errors.New("invalid ")
		}
	}
	return nil
}

// func (rs *regulationStorage) GetOne(id string) *entity.Regulation {
// 	return nil
// }

// func (rs *regulationStorage) GetAll() []*entity.RegulationNamesAndIDsView {
// 	return nil
// }
