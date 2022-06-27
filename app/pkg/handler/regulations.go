package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-b8o/regulations_be/models"
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

// @Summary save doc
// @Tags saveDoc
// @Description save new document
// @ID save-doc
// @Accept json
// @Produce json
// @Param input body models.Doc true "doc"
// @Success 200
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /regulations/save [post]
func (h *Handler) saveDoc(c *gin.Context) {
	doc := &models.Doc{}
	err := c.BindJSON(&doc)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.Writer.WriteHeader(200)
}
