package usecase_regulation

import (
	"context"
	"fmt"
	"prod_serv/internal/domain/entity"
	"strings"
)

type RegulationService interface {
	GetOne(ctx context.Context, regulationId uint64) (entity.Response, entity.Regulation)
	Create(ctx context.Context, regulation entity.Regulation) entity.Response
}
type ChapterService interface {
	GetAllById(ctx context.Context, regulationID uint64) (entity.Response, []*entity.Chapter)
}
type ParagraphService interface {
	GetAllById(ctx context.Context, chapterID uint64) (entity.Response, []entity.Paragraph)
}
type regulationUsecase struct {
	regulationService RegulationService
	chapterService    ChapterService
	paragraphService  ParagraphService
}

func NewRegulationUsecase(regulationService RegulationService, chapterService ChapterService, paragraphService ParagraphService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService, chapterService: chapterService, paragraphService: paragraphService}
}

func (u regulationUsecase) CreateRegulation(ctx context.Context, regulation entity.Regulation) entity.Response {
	return u.regulationService.Create(ctx, regulation)
}

func (u regulationUsecase) GetFullRegulationByID(ctx context.Context, regulationID uint64) (entity.Response, entity.Regulation) {
	resp, regulation := u.regulationService.GetOne(ctx, regulationID)
	respErrors := resp.Errors
	resp, chapters := u.chapterService.GetAllById(ctx, regulationID)
	resp.Errors = append(resp.Errors, respErrors...)

	for _, chapter := range chapters {
		response, paragraphs := u.paragraphService.GetAllById(ctx, chapter.ID)
		if len(response.Errors) > 0 {
			resp.Errors = append(resp.Errors, response.Errors...)
		}

		chapter.Paragraphs = paragraphs
	}

	// fmt.Println(len(chapters[0].Paragraphs))

	regulation.Chapters = chapters

	return resp, regulation
}

func (u regulationUsecase) GetDartFullRegulationByID(ctx context.Context, regulationID uint64) (entity.Response, string) {
	resp, regulation := u.regulationService.GetOne(ctx, regulationID)
	respErrors := resp.Errors
	resp, chapters := u.chapterService.GetAllById(ctx, regulationID)
	resp.Errors = append(resp.Errors, respErrors...)

	dartClass := `
	import 'paragraph.dart';
	import 'chapter.dart';
	
	class Regulation {
		static const int id = %d;
		static const String name = "%s";
		static const String abbreviation = "%s";
		static const List<Chapter> chapters = <Chapter>[
			%s
		];
	}
	`

	chaptersDartString := u.chaptersDart(ctx, chapters)
	return resp, fmt.Sprintf(dartClass, regulationID, regulation.Name, regulation.Abbreviation, chaptersDartString)
}

func (u regulationUsecase) chaptersDart(ctx context.Context, chapters []*entity.Chapter) (dartChaptersString string) {
	dartChapter := `Chapter(id: %d, name: "%s", num: "%s", orderNum: %d , paragraphs: [
		%s
	]),`
	for _, chapter := range chapters {
		_, paragraphs := u.paragraphService.GetAllById(ctx, chapter.ID)

		dartPar := paragraphsDart(paragraphs)

		num := ""
		if len(chapter.Num) > 0 {
			num = chapter.Num
		}
		name := strings.Replace(chapter.Name, "\n", "", -1)
		temp := fmt.Sprintf(dartChapter, chapter.ID, name, strings.TrimSpace(num), chapter.OrderNum, dartPar)
		dartChaptersString += temp

	}
	return dartChaptersString
}

func paragraphsDart(paragraphs []entity.Paragraph) (dartParagraphsList string) {
	for _, p := range paragraphs {
		text := strings.Replace(p.Content, "\n", "", -1)
		dartParagraphsList += fmt.Sprintf(`		Paragraph(id: %d, num: %d, isHTML: %t, paragraphClass: "%s", content: "%s", chapterID: %d),
		`, p.ID, p.Num, p.IsHTML, p.Class, text, p.ChapterID)
	}
	return dartParagraphsList
}

// func (u regulationUsecase) ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView {
// 	return u.regulationService.GetNamesAndIDsOfAllRegulations(ctx)
// }
