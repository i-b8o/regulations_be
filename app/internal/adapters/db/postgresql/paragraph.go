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

func (ps *paragraphStorage) GetAllById(ctx context.Context, chapterID uint64) (entity.Response, []entity.Paragraph) {
	const sql = `SELECT paragraph_id, num, is_html, is_table, class, content, c_id FROM "paragraphs" WHERE c_id = $1 ORDER BY num`
	var resp entity.Response
	var paragraphs []entity.Paragraph

	rows, err := ps.client.Query(ctx, sql, chapterID)
	if err != nil {
		resp.Errors = append(resp.Errors, "Paragraph GetAllByID Query "+err.Error())
		return resp, nil
	}

	defer rows.Close()

	for rows.Next() {
		paragraph := entity.Paragraph{}
		if err = rows.Scan(
			&paragraph.ID, &paragraph.Num, &paragraph.IsHTML, &paragraph.IsTable, &paragraph.Class, &paragraph.Content, &paragraph.ChapterID,
		); err != nil {
			resp.Errors = append(resp.Errors, "Paragraph GetAllByID Next "+err.Error())
			return resp, nil
		}

		paragraphs = append(paragraphs, paragraph)
	}

	return resp, paragraphs

}

func (ps *paragraphStorage) CreateAll(ctx context.Context, paragraphs []entity.Paragraph) entity.Response {
	vals := []interface{}{}
	sql := `INSERT INTO paragraphs ("paragraph_id","num","is_html","is_table","class","content","c_id") VALUES `
	i := 1
	for _, p := range paragraphs {

		sql += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d),", i, i+1, i+2, i+3, i+4, i+5, i+6)
		i = i + 7
		vals = append(vals, p.ID, p.Num, p.IsHTML, p.IsTable, p.Class, p.Content, p.ChapterID)
	}
	sql = sql[:len(sql)-1]
	resp := entity.Response{}

	if _, err := ps.client.Exec(ctx, sql, vals...); err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		return resp
	}

	return resp
}
