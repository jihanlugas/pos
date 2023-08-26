package item

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	testCases := []struct {
		name          string
		id            string
		expectedCode  int
		expectedBody  string
		expectedError bool
	}{
		{
			name:          "Valid ID",
			id:            "1",
			expectedCode:  http.StatusOK,
			expectedBody:  `{"message":"success","data":{"id":"1","name":"example"}}`, // Replace with expected response body
			expectedError: false,
		},
		{
			name:          "Empty ID",
			id:            "",
			expectedCode:  http.StatusBadRequest,
			expectedBody:  `{"message":"data not found"}`, // Replace with expected response body
			expectedError: false,
		},
		{
			name:          "Error from usecase.GetById",
			id:            "123",
			expectedCode:  http.StatusBadRequest,
			expectedBody:  `{"message":"data not found"}`, // Replace with expected response body
			expectedError: true,
		},
	}

	// Create a new echo instance
	e := echo.New()

	// Create a new instance of the Handler
	itemRepo := NewItemRepository()
	itemUsecase := NewItemUsecase(itemRepo)
	itemHandler := ItemHandler(itemUsecase)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request to test the handler
			req := httptest.NewRequest(http.MethodGet, "/item/"+tc.id, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the GetById handler function
			err := itemHandler.GetById(c)

			// Check the error and response code
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedCode, rec.Code)

			// Check the response body
			assert.Equal(t, tc.expectedBody, rec.Body.String())
		})
	}

}

func TestCreate(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := `{"name": "Test Item"}`
	req := httptest.NewRequest(http.MethodPost, "/item", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := Handler{}

	// Test cases
	t.Run("Valid request", func(t *testing.T) {
		// Execute
		err := h.Create(c)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"message":"success","data":{}}`, rec.Body.String())
	})

	t.Run("Invalid request", func(t *testing.T) {
		// Setup
		reqBody := `{"name": ""}`
		req := httptest.NewRequest(http.MethodPost, "/item", strings.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Execute
		err := h.Create(c)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `{"message":"error validation","data":{"name":"name is required"}}`, rec.Body.String())
	})
}
