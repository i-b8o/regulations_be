package storage

import (
	"context"

	"prod_serv/internal/domain/regulation/model"
	db "prod_serv/pkg/client/postgresql/model"

	"prod_serv/pkg/logging"

	sq "github.com/Masterminds/squirrel"
)

type RegulationStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
	logger       *logging.Logger
}

func NewRegulationStorage(client PostgreSQLClient, logger *logging.Logger) RegulationStorage {

	return RegulationStorage{queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), client: client, logger: logger}
}

const (
	scheme = "public"
	table  = "regulations"
)

func (s *RegulationStorage) queryLogger(sql, table string, args []interface{}) *logging.Logger {
	return s.logger.ExtraFields(map[string]interface{}{
		"sql":   sql,
		"table": table,
		"args":  args,
	})
}

func (s *RegulationStorage) All(ctx context.Context) ([]model.Regulation, error) {
	query := s.queryBuilder.Select("regulation_id").
		Column("regulation_name").
		Column("created_at").
		Column("updated_at").From(scheme + "." + table)
	sql, args, err := query.ToSql()
	logger := s.queryLogger(sql, table, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err)
		return nil, err
	}

	logger.Trace("do query")
	rows, err := s.client.Query(ctx, sql, args...)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	list := make([]model.Regulation, 0)

	for rows.Next() {
		r := model.Regulation{}
		if err = rows.Scan(&r.RegulationId, &r.RegulationName, &r.CreatedAt, &r.UpdatedAt); err != nil {
			err = db.ErrScan(err)
			logger.Error(err)
			return nil, err
		}
		list = append(list, r)
	}

	return list, nil
}
