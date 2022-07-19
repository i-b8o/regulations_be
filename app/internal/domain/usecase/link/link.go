package usecase_link

import (
	"context"
	"fmt"
	"prod_serv/internal/domain/entity"
)

type LinkService interface {
	Create(ctx context.Context, link entity.Link) entity.Response
	GetAll(ctx context.Context) (entity.Response, []entity.Link)
	GetAllByChapterID(ctx context.Context, chapterID uint64) (entity.Response, []entity.Link)
}

type ChapterService interface {
	GetOrderNum(ctx context.Context, id uint64) (orderNum uint64, err error)
}

type linkUsecase struct {
	linkService    LinkService
	chapterService ChapterService
}

func NewLinkUsecase(linkService LinkService, chapterService ChapterService) *linkUsecase {
	return &linkUsecase{linkService: linkService, chapterService: chapterService}
}

func (u linkUsecase) CreateLink(ctx context.Context, link entity.Link) entity.Response {
	return u.linkService.Create(ctx, link)
}

func (u linkUsecase) linksDart(ctx context.Context, links []entity.Link) (dartLinksList string) {
	for _, l := range links {
		num, err := u.chapterService.GetOrderNum(ctx, l.ChapterID)
		if err != nil {
			fmt.Println(err)
		}
		dartLinksList += fmt.Sprintf(`		Link(id: %d, chapterNum: %d, ParagraphNum: %d, RID: %d),
		`, l.ID, num, l.ParagraphNum, l.RID)
	}
	return dartLinksList
}

func (u linkUsecase) GetDartAllLinks(ctx context.Context) (entity.Response, string) {
	resp, links := u.linkService.GetAll(ctx)
	respErrors := resp.Errors
	resp.Errors = append(resp.Errors, respErrors...)

	dartClass := `
	import 'link.dart';
	
	class AllLinks {
		static const List<Link> links = <Link>[
			%s
		];
	}
	`

	linkssDartString := u.linksDart(ctx, links)

	return resp, fmt.Sprintf(dartClass, linkssDartString)
}
