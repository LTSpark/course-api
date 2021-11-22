package courses

import (
	"net/http"

	mooc "course-api/internal"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		// Binds request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// Generates uuid for database ID if ID is empty
		if len(req.ID) > 0 {
			uuid, _ := uuid.NewUUID()
			req.ID = uuid.String()
		}

		// Calls mooc dependency
		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.String(http.StatusCreated, "Course created!")
	}
}
