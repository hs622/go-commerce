package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hs622/ecommerce-cart/repository"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils"
)

var CustomErrors = make(map[string]string)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var payload schemas.CreateProductRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {

		var unMarshalType *json.UnmarshalTypeError
		if errors.As(err, &unMarshalType) {

			CustomErrors[strings.ToLower(unMarshalType.Field)] = fmt.Sprintf(
				"field %s must be a %s",
				unMarshalType.Field,
				unMarshalType.Type.String(),
			)

			utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", CustomErrors)
			return
		}

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {

			for _, fieldErr := range validationErrors {
				field := strings.ToLower(fieldErr.Field())
				tag := strings.ToLower(fieldErr.Tag())

				CustomErrors[field] = tag
			}

			utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", CustomErrors)
			return
		}
	}

	if err := h.repo.CreateSingleProduct(ctx.Request.Context(), &payload); err != nil {
		CustomErrors["dbErrors"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create product", CustomErrors)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Product created successfully", payload)
}

// Search Products
//
// @Path_Param ProductId UUID
//
// @Return Product product
func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	var product schemas.CreateProductRequest
	var productId = ctx.Param("productId")

	if !utils.CheckUuid(productId) {
		CustomErrors["Product_Id"] = "invalid product id."
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request id", CustomErrors)
		return
	}

	if err := h.repo.FetchProdcutById(ctx, productId, &product, ctx.Request.URL); err != nil {
		CustomErrors["error"] = "couldn't find the product."
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product Id.", CustomErrors)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Product fetch successfully.", product)
}

// Search Products By Query
//
// @Query limit number
// @Query skip number
// @Query select []string
//
// @Return Products []product | error
func (h *ProductHandler) GetProdcutByQuery(ctx *gin.Context) {
	var products []schemas.CreateProductRequest

	if err := h.repo.FetchProductsWithQuery(ctx, &products, ctx.Request.URL); err != nil {
		log.Fatalln("Inside fetchProductWithQuery: ", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "products fetch successfully.", products)
}

// Update Product
func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var updateProduct schemas.CreateProductRequest
	var product schemas.PatchProductRequest
	var productId = ctx.Param("productId")

	//check for uuid
	if !utils.CheckUuid(productId) {
		CustomErrors["error"] = "couldn't  find the product."
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product Id.", CustomErrors)
		return
	}

	// validation part
	if err := ctx.ShouldBindJSON(&product); err != nil {
		utils.Error(fmt.Sprint(err))

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {

			for _, fieldErr := range validationErrors {
				field := strings.ToLower(fieldErr.Field())
				tag := strings.ToLower(fieldErr.Tag())

				CustomErrors[field] = tag
			}

			utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", CustomErrors)
			return
		}
	}

	// repository method
	if err := h.repo.UpdateSingleProduct(ctx, &product, &productId, &updateProduct); err != nil {
		utils.Error(fmt.Sprint(err))
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Product has been updated sucessfully.", updateProduct)
}

// Delete Products
func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	var result bool = false

	utils.SuccessResponse(ctx, http.StatusOK, "Product has been successfully", result)
}

// Restored Product
func (h *ProductHandler) RestoreProduct(ctx *gin.Context) {
	var result bool = false

	utils.SuccessResponse(ctx, http.StatusOK, "Product has been successfully", result)
}
