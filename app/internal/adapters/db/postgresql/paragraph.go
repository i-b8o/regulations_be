package postgressql

import (
	"context"
	"fmt"
	"prod_serv/internal/domain/entity"
	client "prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient, logger *logging.Logger) *paragraphStorage {
	return &paragraphStorage{client: client}
}

func (cs *paragraphStorage) CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response {
	vals := []interface{}{}
	sql := `INSERT INTO paragraphs ("paragraph_id","num","is_html","class","content","c_id") VALUES `
	i := 1
	for _, p := range paragraphs {

		sql += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d),", i, i+1, i+2, i+3, i+4, i+5)
		i = i + 6
		vals = append(vals, p.ID, p.Num, p.IsHTML, p.Class, p.Content, p.ChapterID)
	}
	sql = sql[:len(sql)-1]
	resp := entity.Response{}

	if _, err := cs.client.Exec(ctx, sql, vals...); err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		return resp
	}

	return resp
}
