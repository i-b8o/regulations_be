package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"prod_serv/internal/controller/http/dto"
	regulation_usecase "prod_serv/internal/domain/usecase/regulation"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	regulationCreate = "/r"
)

type RegulationUsecase interface {
	// ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetRegulationByID(ctx context.Context, id string) *entity.Regulation
	CreateRegulation(ctx context.Context, dto regulation_usecase.CreateRegulationDTO) (uint64, error)
}

type regulationHandler struct {
	regulationUsecase RegulationUsecase
}

func NewRegulationHandler(regulationUsecase RegulationUsecase) *regulationHandler {
	return &regulationHandler{regulationUsecase: regulationUsecase}
}

func (h *regulationHandler) Register(router *httprouter.Router) {
	router.POST(regulationCreate, h.CreateRegulation)
}

// func (h *regulationHandler) ListAllRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	// regulations := h.regulationUsecase.ListAllRegulationNamesAndIDs(context.Background())
// 	w.Write([]byte("regulations"))
// 	w.WriteHeader(http.StatusOK)
// }

func (h *regulationHandler) CreateRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var d dto.CreateRegulationRequestDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return
	}

	// Validation
	if err := d.Validate(); err != nil {
		return
	}
	// MAPPING dto.CreateRegulationRequestDTO --> regulation_usecase.CreateRegulationDTO
	usecaseDTO := regulation_usecase.CreateRegulationDTO{
		RegulationName: d.RegulationName,
	}
	regulationID, err := h.regulationUsecase.CreateRegulation(r.Context(), usecaseDTO)
	if err != nil {
		return
	}
	respDTO := dto.CreateRegulationResponseDTO{
		RegulationID: strconv.FormatUint(regulationID, 10),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respDTO)
}
