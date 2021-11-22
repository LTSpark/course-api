package users

import (
	"errors"
	"net/http"

	mooc "course-api/internal"
	"course-api/internal/creating"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateHandler returns an HTTP handler for Users creation.
func CreateHandler(creatingUserService creating.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		// Binds request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := creatingUserService.CreateUser(ctx, req.ID, req.Name, req.Email, req.Password)
		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidUserID),
				errors.Is(err, mooc.ErrEmptyUserName),
				errors.Is(err, mooc.ErrInvalidUserName),
				errors.Is(err, mooc.ErrEmptyUserPassword),
				errors.Is(err, mooc.ErrInvalidUserEmail):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.String(http.StatusCreated, "User created!")
	}
}
