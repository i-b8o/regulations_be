package handler

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/i-b8o/nonsense"
)

// @Summary regulation
// @Tags regulations
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body pb.CreateUserRequest true "account info"
// @Success 200 {object} pb.CreateUserResponse 1
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /regulations/regulation [post]
func (h *Handler) regulation(c *gin.Context) {
	input := &pb.CreateUserRequest{}
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = mail.ParseAddress(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// Create User
	// TODO separate user exist from error
	resp, err := h.authClient.CreateUser(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if resp.ID == 0 {
		newErrorResponse(c, http.StatusConflict, "")
	}

	// Add confirm token to db
	token := nonsense.RandSeq(100)
	_, err = h.authClient.InsertEmailConfirmToken(c, &pb.InsertEmailConfirmTokenRequest{ID: resp.ID, Token: token})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Send the confirm token to a user email
	r, err := h.mailClient.Confirm(c, &pb.MailConfirmRequest{Url: "bdrop.net/auth/confirmation/" + token, Email: input.Email, Pass: input.Password})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(r.Message) > 0 {
		newErrorResponse(c, http.StatusInternalServerError, r.Message)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": resp.ID})
}
