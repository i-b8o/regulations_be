package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"prod_serv/internal/controller/http/dto"
	usecase "prod_serv/internal/domain/usecase/chapter"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	chapterCreate = "/c"
)

type ChapterUsecase interface {
	// ListAllRegulationNamesAndIDs(ctx context.Context) []*entity.RegulationNamesAndIDsView
	// GetRegulationByID(ctx context.Context, id string) *entity.Regulation
	CreateChapter(ctx context.Context, dto usecase.CreateChapterInput) (usecase.CreateChapterOutput, error)
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

// func (h *chapterHandler) ListAllRegulation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	// regulations := h.regulationUsecase.ListAllRegulationNamesAndIDs(context.Background())
// 	w.Write([]byte("regulations"))
// 	w.WriteHeader(http.StatusOK)
// }

func (h *chapterHandler) CreateChapter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var d dto.CreateChapterRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		// TODO everywhere return err
		return
	}

	// Validation
	if err := d.Validate(); err != nil {
		return
	}
	// MAPPING dto.CreateChapterRequestDTO --> usecase.CreateChapterInput

	usecaseInput := usecase.CreateChapterInput{
		Name:         d.ChapterName,
		Num:          d.ChapterNum,
		RegulationID: d.RegulationID,
	}
	usecaseOutput, err := h.chapterUsecase.CreateChapter(r.Context(), usecaseInput)
	if err != nil {
		return
	}
	respDTO := dto.CreateChapterResponse{
		ChapterID: strconv.FormatUint(usecaseOutput.ChapterID, 10),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respDTO)
}
