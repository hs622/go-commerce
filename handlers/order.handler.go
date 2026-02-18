package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hs622/ecommerce-cart/constants"
	"github.com/hs622/ecommerce-cart/repository"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils"
	"github.com/hs622/ecommerce-cart/utils/validation"
)

type OrderHandler struct {
	repo *repository.OrderRepository
}

func NewOrderHandler(repo *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{
		repo: repo,
	}
}

func (r *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order schemas.CreateOrderRequest
	requestErrors := make(map[string]string)

	if jsonBindingErr := ctx.ShouldBindJSON(&order); jsonBindingErr != nil {

		var unMarshalType *json.UnmarshalTypeError
		if errors.As(jsonBindingErr, &unMarshalType) {
			requestErrors[strings.ToLower(unMarshalType.Field)] = fmt.Sprintf(
				"field %s must be a %s",
				unMarshalType.Field,
				unMarshalType.Type.String(),
			)
		}

		var validationErrors validator.ValidationErrors
		if errors.As(jsonBindingErr, &validationErrors) {
			for _, fe := range jsonBindingErr.(validator.ValidationErrors) {
				requestErrors[utils.ToSnakeCase(strings.Split(fe.Namespace(), ".")[1])] = validation.GetCustomErrorMessage(fe)
			}
		}

		utils.ErrorResponse(ctx, http.StatusBadRequest, string(constants.V_INVALID_ARGUMENT), requestErrors)
		ctx.Abort()
		return
	}

	// call repository function
	if err := r.repo.CreateSingleOrder(ctx.Request.Context(), &order); err != nil {
		utils.Error(fmt.Sprint(err))
		requestErrors["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Unable to create a order.", requestErrors)
		return
	}

	// call response utils function
	utils.SuccessResponse(ctx, http.StatusOK, "Order created successfully.", order)
}

func (r *OrderHandler) UpdateOrder(ctx *gin.Context) {

	var order schemas.CreateOrderRequest
	var payload schemas.PatchOrderRequest
	var productId = ctx.Param("orderId")
	requestErrors := make(map[string]string)

	if !utils.CheckUuid(productId) {
		requestErrors["error"] = "couldn't find the order."
		utils.ErrorResponse(ctx, http.StatusBadRequest, string(constants.V_INVALID_ARGUMENT), requestErrors)
		return
	}

	if jsonBindingErr := ctx.ShouldBindJSON(&payload); jsonBindingErr != nil {

		var unMmarshalType *json.UnmarshalTypeError
		if errors.As(jsonBindingErr, &unMmarshalType) {
			requestErrors[strings.ToLower(unMmarshalType.Field)] = fmt.Sprintf(
				"field %s must be a %s",
				unMmarshalType.Field,
				unMmarshalType.Type.String(),
			)
		}

		var validatorErr validator.ValidationErrors
		if errors.As(jsonBindingErr, &validatorErr) {
			for _, fe := range jsonBindingErr.(validator.ValidationErrors) {
				requestErrors[utils.ToSnakeCase(strings.Split(fe.Namespace(), ".")[1])] = validation.GetCustomErrorMessage(fe)
			}
		}
	}

	if err := r.repo.UpdateSingleOrder(ctx, productId, &payload, &order); err != nil {
		fmt.Println(err)
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Order created successfully.", &order)
}

func (r *OrderHandler) CancelOrder(ctx *gin.Context) {

	orderId := ctx.Param("orderId")
	requestError := make(map[string]string)

	if !utils.CheckUuid(orderId) {
		requestError["error"] = "Invalid order id."
		utils.ErrorResponse(ctx, http.StatusBadRequest, string(constants.V_INVALID_ARGUMENT), requestError)
		return
	}

	if err := r.repo.SuspendOrder(ctx, orderId); err != nil {
		requestError["s"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, string(constants.ERROR_M_UNABLE_CANCEL), requestError)
		return
	}

	utils.SuccessResponse(ctx, http.StatusBadRequest, "Order cancelled.", nil)
}

func (r *OrderHandler) SoftDeleteOrder(ctx *gin.Context) {

	var order schemas.CreateOrderRequest
	orderId := ctx.Param("orderId")
	requestError := make(map[string]string)

	if !utils.CheckUuid(orderId) {
		requestError["error"] = "Invalid order id."
		utils.ErrorResponse(ctx, http.StatusBadRequest, string(constants.V_INVALID_ARGUMENT), requestError)
		return
	}

	if err := r.repo.SoftDeleteOrder(ctx, orderId, &order); err != nil {
		fmt.Println(err)
		requestError["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, string(constants.ERROR_M_UNABLE_DELETE), requestError)
		return
	}

	status := &order.Status
	utils.SuccessResponse(ctx, http.StatusOK, "Order deleted successfully.", status)
}

func (r *OrderHandler) RestoreOrder(ctx *gin.Context) {}

func (r *OrderHandler) GetOrders(ctx *gin.Context) {}

func (r *OrderHandler) GetOrder(ctx *gin.Context) {}
