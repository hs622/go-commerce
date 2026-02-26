package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/repository"
	"github.com/hs622/ecommerce-cart/utils"
)

type PaymentHandler struct {
	repo *repository.PaymentRepository
}

func NewPaymentHandler(repo *repository.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{
		repo: repo,
	}
}

func (h *PaymentHandler) CreatePaymentIntent(ctx *gin.Context) {

	requestError := make(map[string]string)

	p, err := h.repo.CreatePayment(ctx.Request.Context())
	if err != nil {
		requestError["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Bad request.", requestError)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "worked!", p)
}

func (h *PaymentHandler) UpdatePaymentIntent(ctx *gin.Context) {

	intentId := ctx.Param("paymentId")
	requestError := make(map[string]string)

	p, err := h.repo.UpdatePayment(ctx.Request.Context(), intentId)
	if err != nil {
		fmt.Println("requestError")
		requestError["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request.", requestError)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Payment updated successfully.", p)
}

func (r *PaymentHandler) DeletePaymentIntent(ctx *gin.Context) {}
func (r *PaymentHandler) GetPaymentIntents(ctx *gin.Context) {

}
func (h *PaymentHandler) GetPaymentIntent(ctx *gin.Context) {

	requestError := make(map[string]string)

	p, err := h.repo.Payment(ctx.Request.Context(), ctx.Param("paymentId"))
	if err != nil {
		requestError["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request!", requestError)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Fetched!", p)
}
