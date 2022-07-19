package v1

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"prod_serv/internal/controller/http/dto"
// 	"prod_serv/internal/domain/entity"

// 	"github.com/julienschmidt/httprouter"
// )

// const (
// // linksDart = "/lfdart"
// )

// type LinkUsecase interface {
// 	GetDartAllLinks(ctx context.Context) (entity.Response, string)
// }

// type linkHandler struct {
// 	linkUsecase LinkUsecase
// }

// func NewLinkHandler(regulationUsecase LinkUsecase) *linkHandler {
// 	return &linkHandler{linkUsecase: regulationUsecase}
// }

// func (h *linkHandler) Register(router *httprouter.Router) {
// 	router.POST(linksDart, h.AllLinksDart)
// }

// func (h *linkHandler) AllLinksDart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	// Set headers
// 	w.Header().Set("Content-Type", "application/json")

// 	// Input and Output
// 	var input dto.GetAllLinksRequestDTO
// 	var out dto.GetAllLinksDartResponseDTO

// 	// Get JSON request
// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		fmt.Println(err)
// 		json.NewEncoder(w).Encode(out)
// 		return
// 	}
// 	defer r.Body.Close()

// 	// MAPPING dto.CreateRegulationRequestDTO --> string

// 	// Usecase
// 	_, dartStr := h.linkUsecase.GetDartAllLinks(r.Context())

// 	// MAPPING entity.Regulation --> dto.GetFullRegulationResponseDTO

// 	// start := `List<Chapter> allChapters = <Chapter>[`
// 	// str := `
// 	// `

// 	// for _, chapter := range regulation.Chapters {
// 	// 	str += fmt.Sprintf(`Chapter(
// 	// 		ID: "%d";
// 	// 		Name: "%s";
// 	// 		Num: "%s";
// 	// 		`, chapter.ID, chapter.Name, chapter.Num)
// 	// }

// 	// dart := start + str

// 	w.Write([]byte(dartStr))

// }
