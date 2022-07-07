package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"prod_serv/internal/controller/http/dto"
	"prod_serv/internal/domain/entity"
	usecase_paragraph "prod_serv/internal/domain/usecase/paragraph"

	"github.com/julienschmidt/httprouter"
)

const (
	paragraphsCreate = "/p"
)

type ParagraphUsecase interface {
	// ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetRegulationByID(ctx context.Context, id string) *entity.Regulation
	CreateParagraphs(ctx context.Context, dto usecase_paragraph.CreateParagraphsInput) usecase_paragraph.CreateParagraphsOutput
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
	var d dto.CreateParagraphsRequest
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		resp := dto.ErrorResponse{Message: "Failed to decode"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var usecaseIn usecase_paragraph.CreateParagraphsInput
	// MAPPING dto.CreateParagraphsRequest --> usecase.CreateParagraphsInput
	for _, p := range d.Paragraphs {
		// Validation
		if err := p.Validate(); err != nil {
			resp := dto.ErrorResponse{Message: err.Error()}
			json.NewEncoder(w).Encode(resp)
			return
		}

		paragraph := entity.Paragraph{
			ID:        p.ParagraphID,
			Num:       p.ParagraphOrderNum,
			Class:     p.ParagraphClass,
			Content:   p.ParagraphText,
			ChapterID: p.ChapterID,
		}
		usecaseIn.Paragraphs = append(usecaseIn.Paragraphs, paragraph)
	}

	usecaseOutput := h.paragraphUsecase.CreateParagraphs(r.Context(), usecaseIn)

	resp := dto.CreateParagraphsResponse{
		Message: usecaseOutput.Message,
	}
	json.NewEncoder(w).Encode(resp)
}
