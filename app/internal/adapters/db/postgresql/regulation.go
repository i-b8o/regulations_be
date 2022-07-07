package postgressql

import (
	"context"
	"errors"
	"log"
	dto "prod_serv/internal/adapters/db/dto"
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

func (rs *regulationStorage) Create(ctx context.Context, params dto.CreateRegulationInput) (dto.CreateRegulationOutput, error) {
	t := time.Now()

	const sql = `INSERT INTO regulations ("name", "abbreviation", "created_at") VALUES ($1, $2, $3) RETURNING "id"`

	row := rs.client.QueryRow(ctx, sql, params.RegulationName, params.Abbreviation, t)
	var regulationID uint64
	resp := dto.CreateRegulationOutput{}
	switch err := row.Scan(&regulationID); {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		return resp, err
	case err != nil:
		if sqlErr := rs.regulationPgError(err); sqlErr != nil {
			return resp, sqlErr
		}
		log.Printf("cannot create regulation on database: %v\n", err)
		return resp, errors.New("cannot create regulation on database")
	}
	resp.RegulationID = regulationID
	return resp, nil
}

func (rs *regulationStorage) regulationPgError(err error) error {
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
