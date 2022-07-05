package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"prod_serv/internal/controller/http/dto"
	"prod_serv/internal/domain/entity"
	regulation_usecase "prod_serv/internal/domain/usecase/regulation"

	"github.com/julienschmidt/httprouter"
)

const (
	regulationURL  = "/rs/:r_id"
	regulationsURL = "/rs"
)

type RegulationUsecase interface {
	ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView
	GetRegulationByID(ctx context.Context, id string) *entity.Regulation
	CreateRegulation(ctx context.Context, reg regulation_usecase.CreateRegulationDTO) error
}

type regulationHandler struct {
	regulationUsecase RegulationUsecase
}

func NewRegulationHandler(regulationUsecase RegulationUsecase) *regulationHandler {
	return &regulationHandler{regulationUsecase: regulationUsecase}
}

func (h *regulationHandler) Register(router *httprouter.Router) {
	router.GET(regulationURL, h.ListAllRegulation)
}

func (h *regulationHandler) ListAllRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// regulations := h.regulationUsecase.ListAllRegulationNamesAndIDs(context.Background())
	w.Write([]byte("regulations"))
	w.WriteHeader(http.StatusOK)
}

func (h *regulationHandler) CreateRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var d dto.CreateRegulationDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return
	}

	//TODO validate

	// MAPPING dto.CreateBookDTO --> regulation_usecase.CreateRegulationDTO
	usecaseDTO := regulation_usecase.CreateRegulationDTO{
		RegulationName: d.RegulationName,
	}
	err := h.regulationUsecase.CreateRegulation(r.Context(), usecaseDTO)
	if err != nil {
		return
	}
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}
