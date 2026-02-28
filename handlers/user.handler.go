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

	utils.SuccessResponse(ctx, http.StatusOK, "user created successfully.", &user)
}

func (r *UserHandler) UpdateUser(ctx *gin.Context)  {}
func (r *UserHandler) SuspendUser(ctx *gin.Context) {}
func (r *UserHandler) DeleteUser(ctx *gin.Context)  {}
func (r *UserHandler) GetUser(ctx *gin.Context)     {}
func (r *UserHandler) GetUsers(ctx *gin.Context)    {}
