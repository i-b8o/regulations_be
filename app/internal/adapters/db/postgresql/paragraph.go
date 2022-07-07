package postgressql

import (
	"context"
	"fmt"
	"prod_serv/internal/adapters/db/dto"
	client "prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient, logger *logging.Logger) *paragraphStorage {
	return &paragraphStorage{client: client}
}

func (cs *paragraphStorage) CreateAll(ctx context.Context, in dto.CreateParagraphsInput) dto.CreateParagraphsOutput {
	vals := []interface{}{}
	sql := `INSERT INTO paragraphs ("num","class","content","c_id") VALUES `
	i := 1
	for _, p := range in.Paragraphs {

		sql += fmt.Sprintf("($%d, $%d, $%d, $%d),", i, i+1, i+2, i+3)
		i = i + 5
		vals = append(vals, p.Num, p.Class, p.Content, p.ChapterID)
	}
	sql = sql[:len(sql)-1]
	out := dto.CreateParagraphsOutput{}

	if _, err := cs.client.Exec(ctx, sql, vals...); err != nil {
		out.Message = "SQL exec error " + err.Error()
		return out
	}

	return out
}
