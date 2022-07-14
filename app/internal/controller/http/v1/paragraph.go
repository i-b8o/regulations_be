package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"prod_serv/internal/controller/http/dto"
	"prod_serv/internal/domain/entity"

	"github.com/julienschmidt/httprouter"
)

const (
	paragraphsCreate = "/p"
)

type ParagraphUsecase interface {
	// ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetRegulationByID(ctx context.Context, id string) *entity.Regulation
	CreateParagraphs(ctx context.Context, paragraphs []entity.Paragraph) entity.Response
}

type paragraphHandler struct {
	paragraphUsecase ParagraphUsecase
}

func NewParagraphHandler(paragraphUsecase ParagraphUsecase) *paragraphHandler {
	return &paragraphHandler{paragraphUsecase: paragraphUsecase}
}

func (h *paragraphHandler) Register(router *httprouter.Router) {
	router.POST(paragraphsCreate, h.CreateParagraphs)
}

func (h *paragraphHandler) CreateParagraphs(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// Input and Output
	var d dto.CreateParagraphsRequest
	var response entity.Response
	var paragraphs []entity.Paragraph

	// Get JSON request
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println(err)
		response.Errors = append(response.Errors, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	defer r.Body.Close()

	// MAPPING dto.CreateParagraphsRequest --> []entity.Paragraph
	for _, p := range d.Paragraphs {
		fmt.Println("a ", p.IsHTML)
		// Validation
		if s, err := p.Validate(); err != nil {
			fmt.Println("Warn", s, "Err", err)
			if s != "" {
				response.Warnings = append(response.Warnings, s)
			}

			response.Errors = append(response.Errors, err.Error())
			json.NewEncoder(w).Encode(response)
			return
		}

		paragraph := entity.Paragraph{
			ID:        p.ParagraphID,
			Num:       p.ParagraphOrderNum,
			IsHTML:    p.IsHTML,
			IsTable:   p.IsTable,
			Class:     p.ParagraphClass,
			Content:   p.ParagraphText,
			ChapterID: p.ChapterID,
		}
		paragraphs = append(paragraphs, paragraph)
	}

	// Usecase
	response = h.paragraphUsecase.CreateParagraphs(r.Context(), paragraphs)

	// Response
	json.NewEncoder(w).Encode(response)
}
