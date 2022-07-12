package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"prod_serv/internal/controller/http/dto"
	"prod_serv/internal/domain/entity"

	"github.com/julienschmidt/httprouter"
)

const (
	chapterCreate = "/c"
)

type ChapterUsecase interface {
	CreateChapter(ctx context.Context, chapter entity.Chapter) entity.Response
}

type chapterHandler struct {
	chapterUsecase ChapterUsecase
}

func NewChapterHandler(regulationUsecase ChapterUsecase) *chapterHandler {
	return &chapterHandler{chapterUsecase: regulationUsecase}
}

func (h *chapterHandler) Register(router *httprouter.Router) {
	router.POST(chapterCreate, h.CreateChapter)
}

func (h *chapterHandler) CreateChapter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// Input and Output
	var d dto.CreateChapterRequest
	var response entity.Response

	// Get JSON request
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		response.Errors = append(response.Errors, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	defer r.Body.Close()

	// Validation
	if s, err := d.Validate(); err != nil {
		response.Errors = append(response.Errors, err.Error())

		if s != "" {
			response.Warnings = append(response.Warnings, s)
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// MAPPING dto.CreateChapterRequestDTO --> entity.Chapter
	chapter := entity.Chapter{
		Name:         d.ChapterName,
		Num:          d.ChapterNum,
		RegulationID: d.RegulationID,
		OrderNum:     d.OrderNum,
	}

	// Usecase
	response = h.chapterUsecase.CreateChapter(r.Context(), chapter)

	// Response
	json.NewEncoder(w).Encode(response)
}
