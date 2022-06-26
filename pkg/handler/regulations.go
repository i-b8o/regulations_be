package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary get doc
// @Tags getDoc
// @Description return regulation doc
// @ID get-doc
// @Accept json
// @Produce json
// @Param abbreviation query string true "Regulation Abbreviation"
// @Success 200 {object} models.Doc
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /regulations/{abbreviation} [get]
func (h *Handler) getDoc(c *gin.Context) {
	abbreviation := c.Param("abbreviation")
	fmt.Println(abbreviation)
}
