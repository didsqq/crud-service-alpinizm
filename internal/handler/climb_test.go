package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/service"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClimbService - мок для сервиса восхождений
type MockClimbService struct {
	mock.Mock
}

func (m *MockClimbService) GetAll(ctx context.Context, mountainID, categoryID int) ([]domain.Climb, error) {
	args := m.Called(ctx, mountainID, categoryID)
	return args.Get(0).([]domain.Climb), args.Error(1)
}

func (m *MockClimbService) GetById(ctx context.Context, climbID int64) (domain.Climb, error) {
	args := m.Called(ctx, climbID)
	return args.Get(0).(domain.Climb), args.Error(1)
}

func (m *MockClimbService) RecordAlpinistClimb(ctx context.Context, alpinistID, climbID int64) error {
	args := m.Called(ctx, alpinistID, climbID)
	return args.Error(0)
}

func (m *MockClimbService) GetAlpinistClimb(ctx context.Context, alpinistID int64) ([]domain.Climb, error) {
	args := m.Called(ctx, alpinistID)
	return args.Get(0).([]domain.Climb), args.Error(1)
}

func (m *MockClimbService) GetAllCategoryOfDifficulty(ctx context.Context) ([]domain.CategoryOfDifficulty, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.CategoryOfDifficulty), args.Error(1)
}

func TestHandler_GetAllClimbs(t *testing.T) {
	tests := []struct {
		name           string
		mountainID     string
		categoryID     string
		mockSetup      func(*MockClimbService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:       "успешное получение всех восхождений",
			mountainID: "",
			categoryID: "",
			mockSetup: func(m *MockClimbService) {
				m.On("GetAll", mock.Anything, 0, 0).Return([]domain.Climb{
					{ID: 1, Title: "Восхождение 1"},
					{ID: 2, Title: "Восхождение 2"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.Climb{
				{ID: 1, Title: "Восхождение 1"},
				{ID: 2, Title: "Восхождение 2"},
			},
		},
		{
			name:       "фильтрация по горе",
			mountainID: "1",
			categoryID: "",
			mockSetup: func(m *MockClimbService) {
				m.On("GetAll", mock.Anything, 1, 0).Return([]domain.Climb{
					{ID: 1, Title: "Восхождение на гору 1"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.Climb{
				{ID: 1, Title: "Восхождение на гору 1"},
			},
		},
		{
			name:       "фильтрация по категории",
			mountainID: "",
			categoryID: "2",
			mockSetup: func(m *MockClimbService) {
				m.On("GetAll", mock.Anything, 0, 2).Return([]domain.Climb{
					{ID: 2, Title: "Восхождение категории 2"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.Climb{
				{ID: 2, Title: "Восхождение категории 2"},
			},
		},
		{
			name:           "неверный mountainId",
			mountainID:     "abc",
			categoryID:     "",
			mockSetup:      func(m *MockClimbService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Неверный mountainId",
			},
		},
		{
			name:           "неверный categoryId",
			mountainID:     "",
			categoryID:     "xyz",
			mockSetup:      func(m *MockClimbService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Неверный categoryId",
			},
		},
		{
			name:       "ошибка получения восхождений",
			mountainID: "",
			categoryID: "",
			mockSetup: func(m *MockClimbService) {
				m.On("GetAll", mock.Anything, 0, 0).Return([]domain.Climb{}, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения восхождений",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClimbService := new(MockClimbService)
			tt.mockSetup(mockClimbService)

			h := &Handler{
				services: &service.Service{
					Climbs: mockClimbService,
				},
			}

			url := "/api/climb?"
			if tt.mountainID != "" {
				url += "mountainId=" + tt.mountainID + "&"
			}
			if tt.categoryID != "" {
				url += "categoryId=" + tt.categoryID + "&"
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			h.getAllClimbs(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			expectedJSON, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)
			var expectedResponse interface{}
			err = json.Unmarshal(expectedJSON, &expectedResponse)
			assert.NoError(t, err)

			assert.Equal(t, expectedResponse, response)
			mockClimbService.AssertExpectations(t)
		})
	}
}

func TestHandler_GetClimb(t *testing.T) {
	tests := []struct {
		name           string
		climbID        string
		mockSetup      func(*MockClimbService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:    "успешное получение восхождения",
			climbID: "1",
			mockSetup: func(m *MockClimbService) {
				m.On("GetById", mock.Anything, int64(1)).Return(domain.Climb{
					ID:    1,
					Title: "Восхождение на Эльбрус",
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: &domain.Climb{
				ID:    1,
				Title: "Восхождение на Эльбрус",
			},
		},
		{
			name:           "неверный ID",
			climbID:        "abc",
			mockSetup:      func(m *MockClimbService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Неверный climbId",
			},
		},
		{
			name:    "ошибка получения восхождения",
			climbID: "999",
			mockSetup: func(m *MockClimbService) {
				m.On("GetById", mock.Anything, int64(999)).Return(domain.Climb{}, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения восхождения",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClimbService := new(MockClimbService)
			tt.mockSetup(mockClimbService)

			h := &Handler{
				services: &service.Service{
					Climbs: mockClimbService,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/api/climb/"+tt.climbID, nil)

			// Создаем контекст с параметром chi
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.climbID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()

			h.getClimb(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			expectedJSON, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)
			var expectedResponse interface{}
			err = json.Unmarshal(expectedJSON, &expectedResponse)
			assert.NoError(t, err)

			assert.Equal(t, expectedResponse, response)
			mockClimbService.AssertExpectations(t)
		})
	}
}
