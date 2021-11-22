package users

import (
	"bytes"
	"course-api/internal/creating"
	"course-api/internal/platform/storage/storagemocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/huandu/go-assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {

	// Apply mocks
	userRepository := new(storagemocks.UserRepository)
	userRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.User")).Return(nil)

	createUserSrv := creating.NewUserService(userRepository)

	// Create gin server on test mode
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Register test routes
	r.POST("/users", CreateHandler(createUserSrv))

	t.Run("given an invalid id request it should return 400", func(t *testing.T) {
		createRequest := createRequest{
			ID:       "invalid-id",
			Name:     "First Test User",
			Email:    "test@email.com",
			Password: "contraseña",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given a long name request it should return 400", func(t *testing.T) {
		createRequest := createRequest{
			ID:       "6b0d0378-4b50-11ec-81d3-0242ac130002",
			Name:     "namethatismorethanthirtycharacterslongandmore",
			Email:    "test@email.com",
			Password: "contraseña",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given an invalid email request it should return 400", func(t *testing.T) {
		createRequest := createRequest{
			ID:       "6b0d0378-4b50-11ec-81d3-0242ac130002",
			Name:     "Test User",
			Email:    "this is not an email",
			Password: "contraseña",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given an empty password request it should return 400", func(t *testing.T) {
		createRequest := createRequest{
			ID:       "6b0d0378-4b50-11ec-81d3-0242ac130002",
			Name:     "Test User",
			Email:    "test@email.com",
			Password: "",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given a valid request it should return 201", func(t *testing.T) {
		createRequest := createRequest{
			ID:       "6b0d0378-4b50-11ec-81d3-0242ac130002",
			Name:     "Valid User",
			Email:    "email@email.com",
			Password: "contraseña",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)

	})

}
