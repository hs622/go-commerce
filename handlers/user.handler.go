package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hs622/ecommerce-cart/repository"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils"
	"github.com/hs622/ecommerce-cart/utils/database"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (r *UserHandler) CreateUser(ctx *gin.Context) {

	requestError := make(map[string]string)
	var user schemas.CreateUserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {

		var unMarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unMarshalErr) {
			requestError[strings.ToLower(unMarshalErr.Field)] = fmt.Sprintf(
				"field %s must be a %s",
				unMarshalErr.Field,
				unMarshalErr.Type.String(),
			)
		}

		var validate validator.ValidationErrors
		if errors.As(err, &validate) {
			for _, f := range err.(validator.ValidationErrors) {
				requestError[strings.ToLower(f.Field())] = strings.ToLower(f.Tag())
			}
		}

		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", requestError)
		return
	}

	if err := r.repo.New(ctx, &user); err != nil {
		requestError["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Unable to create user.", requestError)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "user created successfully.", &user)
}

func (r *UserHandler) UpdateUser(ctx *gin.Context) {

	userId := ctx.Param("userId")
	requestError := make(map[string]string)
	var fields schemas.PatchUserRequest
	var user schemas.CreateUserRequest

	if !utils.CheckUuid(userId) {
		requestError["error"] = "invalid user id."
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid resource id.", requestError)
		return
	}

	if err := ctx.ShouldBindJSON(&fields); err != nil {

		var unMarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unMarshalErr) {
			requestError[strings.ToLower(unMarshalErr.Field)] = fmt.Sprintf(
				"field %s must be a %s",
				unMarshalErr.Field,
				unMarshalErr.Type.String(),
			)
		}

		var validate validator.ValidationErrors
		if errors.As(err, &validate) {
			for _, f := range err.(validator.ValidationErrors) {
				requestError[strings.ToLower(f.Field())] = strings.ToLower(f.Tag())
			}
		}

		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload.", requestError)
		return
	}

	if err := r.repo.UpdateUser(ctx.Request.Context(), userId, &fields, &user); err != nil {
		requestError["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Unable to update the field.", requestError)
		return
	}

	utils.SuccessResponse(ctx, http.StatusAccepted, "User update successfully.", &user)
}

func (r *UserHandler) GetUser(ctx *gin.Context) {

	var user schemas.CreateUserRequest
	requestErrors := make(map[string]string)

	if err := r.repo.FetchUser(
		ctx,
		&user,
		database.UserFetchOptions{
			Url:      ctx.Request.URL,
			FullPath: ctx.FullPath(),
		}); err != nil {
		fmt.Println(err)
		requestErrors["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request", requestErrors)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Fetch user successfully.", &user)
}

func (r *UserHandler) GetUsers(ctx *gin.Context) {

	var users []schemas.CreateUserRequest
	requestErrors := make(map[string]string)

	if err := r.repo.FetchUsers(ctx, ctx.Request.URL, &users); err != nil {
		requestErrors["error"] = err.Error()
		utils.ErrorResponse(ctx, http.StatusBadRequest, "", requestErrors)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "fetch user successful.", users)
}
