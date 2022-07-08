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
	regulationCreate  = "/r"
	regulationGetFull = "/rgf"
)

type RegulationUsecase interface {
	CreateRegulation(ctx context.Context, regulation entity.Regulation) entity.Response
	GetFullRegulationByID(ctx context.Context, regulationID uint64) (entity.Response, entity.Regulation)
}

type regulationHandler struct {
	regulationUsecase RegulationUsecase
}

func NewRegulationHandler(regulationUsecase RegulationUsecase) *regulationHandler {
	return &regulationHandler{regulationUsecase: regulationUsecase}
}

func (h *regulationHandler) Register(router *httprouter.Router) {
	router.POST(regulationCreate, h.CreateRegulation)
	router.POST(regulationGetFull, h.GetFullRegulation)
}

func (h *regulationHandler) GetFullRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// Input and Output
	var input dto.GetFullRegulationRequestDTO
	var out dto.GetFullRegulationResponseDTO

	// Get JSON request
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		out.Response.Errors = append(out.Response.Errors, err.Error())
		json.NewEncoder(w).Encode(out)
		return
	}
	defer r.Body.Close()

	// MAPPING dto.CreateRegulationRequestDTO --> string
	usecaseRegulationID := input.RegulationID

	// Usecase
	response, regulation := h.regulationUsecase.GetFullRegulationByID(r.Context(), usecaseRegulationID)

	// MAPPING entity.Regulation --> dto.GetFullRegulationResponseDTO
	out.Response = response
	out.Regulation = regulation
	// out.Response = response
	// out.RegulationID = regulation.Id
	// out.Abbreviation = regulation.Abbreviation
	// out.RegulationName = regulation.Name
	// var chapters []dto.Chapter
	// for _, c := range regulation.Chapters {
	// 	var chapter dto.Chapter
	// 	chapter.ID = c.ID
	// 	chapter.Name = c.Name
	// 	chapter.Num = c.Num
	// 	// chapter.Paragraphs = c.Paragraphs
	// 	chapters = append(chapters, chapter)
	// }
	// out.Chapters = chapters
	// Response
	json.NewEncoder(w).Encode(out)

}

func (h *regulationHandler) CreateRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// Input and Output
	var input dto.CreateRegulationRequestDTO
	var response entity.Response

	// Get JSON request
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Errors = append(response.Errors, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	defer r.Body.Close()

	// Validation
	if err := input.Validate(); err != nil {
		response.Errors = append(response.Errors, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}

	// MAPPING dto.CreateRegulationRequestDTO --> entity.Regulation
	regulation := entity.Regulation{
		Name:         input.RegulationName,
		Abbreviation: input.Abbreviation,
	}

	// Usecase
	response = h.regulationUsecase.CreateRegulation(r.Context(), regulation)

	// Response
	json.NewEncoder(w).Encode(response)
}
