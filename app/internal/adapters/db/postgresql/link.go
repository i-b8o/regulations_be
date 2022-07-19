package postgressql

import (
	"context"
	"fmt"
	"prod_serv/internal/domain/entity"
	client "prod_serv/pkg/client/postgresql"
	"prod_serv/pkg/logging"
)

type linkStorage struct {
	client client.PostgreSQLClient
}

func NewLinkStorage(client client.PostgreSQLClient, logger *logging.Logger) *linkStorage {
	return &linkStorage{client: client}
}

func (ps *linkStorage) GetAll(ctx context.Context) (entity.Response, []*entity.Link) {
	const sql = `SELECT id,c_id,paragraph_num FROM "links" ORDER BY c_id`
	var resp entity.Response
	var links []*entity.Link

	rows, err := ps.client.Query(ctx, sql)
	if err != nil {
		resp.Errors = append(resp.Errors, "Paragraph GetAllByID Query "+err.Error())
		return resp, nil
	}

	defer rows.Close()

	for rows.Next() {
		link := &entity.Link{}
		if err = rows.Scan(
			&link.ID, &link.ChapterID, &link.ParagraphNum); err != nil {
			resp.Errors = append(resp.Errors, "Link GetAll Next "+err.Error())
			return resp, nil
		}

		links = append(links, link)
	}

	return resp, links
}

func (ps *linkStorage) GetAllByChapterID(ctx context.Context, chapterID uint64) (entity.Response, []*entity.Link) {
	const sql = `SELECT id,c_id,paragraph_num FROM "links" WHERE c_id = $1 ORDER BY c_id`
	var resp entity.Response
	var links []*entity.Link
	rows, err := ps.client.Query(ctx, sql, chapterID)

	if err != nil {
		resp.Errors = append(resp.Errors, "Paragraph GetAllByID Query "+err.Error())
		fmt.Println(err)
		return resp, nil
	}

	defer rows.Close()

	for rows.Next() {
		link := &entity.Link{}
		if err = rows.Scan(
			&link.ID, &link.ChapterID, &link.ParagraphNum); err != nil {
			resp.Errors = append(resp.Errors, "Link GetAll Next "+err.Error())
			return resp, nil
		}

		links = append(links, link)
	}

	return resp, links
}

func (ps *linkStorage) Create(ctx context.Context, link entity.Link) entity.Response {
	sql := `INSERT INTO links ("id","c_id","paragraph_num") VALUES ($1,$2,$3)`

	resp := entity.Response{}

	if _, err := ps.client.Exec(ctx, sql, link.ID, link.ChapterID, link.ParagraphNum); err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		return resp
	}

	return resp
}
