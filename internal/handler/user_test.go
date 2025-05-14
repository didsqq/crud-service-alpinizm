package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
	"github.com/didsqq/crud-service-alpinizm/internal/service"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Login(ctx context.Context, username, password string) (string, error) {
	args := m.Called(ctx, username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) Create(ctx context.Context, user domain.User) (int, error) {
	args := m.Called(ctx, user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserService) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserService) GetAllSportCategory(ctx context.Context) ([]domain.SportCategory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.SportCategory), args.Error(1)
}

func (m *MockUserService) CheckToken(ctx context.Context, token string) (bool, error) {
	args := m.Called(ctx, token)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserService) GetAllAlpinists(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserService) GetAll(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserService) CancelAlpinistClimb(ctx context.Context, alpinistID int64, climbID int64) error {
	args := m.Called(ctx, alpinistID, climbID)
	return args.Error(0)
}

func (m *MockUserService) DeleteAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	args := m.Called(ctx, alpinistID, equipmentID)
	return args.Error(0)
}

func TestHandler_LoginUser(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "успешный вход пользователя",
			requestBody: domain.User{
				Username: "testuser",
				Password: "password123",
			},
			mockSetup: func(m *MockUserService) {
				m.On("Login", mock.Anything, "testuser", "password123").Return("jwt-token-123", nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusOK,
				Type:    "success",
				Message: "jwt-token-123",
			},
		},
		{
			name: "пользователь не найден",
			requestBody: domain.User{
				Username: "wronguser",
				Password: "password123",
			},
			mockSetup: func(m *MockUserService) {
				m.On("Login", mock.Anything, "wronguser", "password123").Return("", repository.ErrUserNotFound)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Пользователь не найден",
			},
		},
		{
			name: "ошибка сервера при входе",
			requestBody: domain.User{
				Username: "testuser",
				Password: "password123",
			},
			mockSetup: func(m *MockUserService) {
				m.On("Login", mock.Anything, "testuser", "password123").Return("", errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка входа пользователя",
			},
		},
		{
			name:           "некорректный JSON",
			requestBody:    "invalid json",
			mockSetup:      func(m *MockUserService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Ошибка парсинга json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockUserService)
			tt.mockSetup(mockUserService)

			h := &Handler{
				services: &service.Service{
					User: mockUserService,
				},
			}

			var bodyBytes []byte
			var err error
			if tt.requestBody == "invalid json" {
				bodyBytes = []byte("invalid json")
			} else {
				bodyBytes, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.loginUser(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err = json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			expectedJSON, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)
			var expectedResponse interface{}
			err = json.Unmarshal(expectedJSON, &expectedResponse)
			assert.NoError(t, err)

			assert.Equal(t, expectedResponse, response)
			mockUserService.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateUser(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "успешное создание пользователя",
			requestBody: domain.User{
				Username:        "testuser",
				Password:        "password123",
				Name:            "Иван",
				Surname:         "Петров",
				Address:         "ул. Ленина, 10",
				Phone:           "+7-900-123-45-67",
				Sex:             "мужской",
				IdSportCategory: 1,
			},
			mockSetup: func(m *MockUserService) {
				m.On("Create", mock.Anything, mock.AnythingOfType("domain.User")).Return(1, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusOK,
				Type:    "success",
				Message: "Пользователь создан c id: 1",
			},
		},
		{
			name: "username уже занят",
			requestBody: domain.User{
				Username:        "existinguser",
				Password:        "password123",
				Name:            "Иван",
				Surname:         "Петров",
				Address:         "ул. Ленина, 10",
				Phone:           "+7-900-123-45-67",
				Sex:             "мужской",
				IdSportCategory: 1,
			},
			mockSetup: func(m *MockUserService) {
				m.On("Create", mock.Anything, mock.AnythingOfType("domain.User")).Return(0, repository.ErrUserNameExist)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Username занят",
			},
		},
		{
			name: "ошибка создания пользователя",
			requestBody: domain.User{
				Username:        "testuser",
				Password:        "password123",
				Name:            "Иван",
				Surname:         "Петров",
				Address:         "ул. Ленина, 10",
				Phone:           "+7-900-123-45-67",
				Sex:             "мужской",
				IdSportCategory: 1,
			},
			mockSetup: func(m *MockUserService) {
				m.On("Create", mock.Anything, mock.AnythingOfType("domain.User")).Return(0, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка создания пользователя",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockUserService)
			tt.mockSetup(mockUserService)

			h := &Handler{
				services: &service.Service{
					User: mockUserService,
				},
			}

			bodyBytes, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/user/registration", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.createUser(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response interface{}
			err = json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			expectedJSON, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)
			var expectedResponse interface{}
			err = json.Unmarshal(expectedJSON, &expectedResponse)
			assert.NoError(t, err)

			assert.Equal(t, expectedResponse, response)
			mockUserService.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:   "успешное удаление пользователя",
			userID: "1",
			mockSetup: func(m *MockUserService) {
				m.On("Delete", mock.Anything, 1).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusOK,
				Type:    "success",
				Message: "Пользователь успешно удалён",
			},
		},
		{
			name:   "пользователь не найден",
			userID: "999",
			mockSetup: func(m *MockUserService) {
				m.On("Delete", mock.Anything, 999).Return(repository.ErrUserNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusNotFound,
				Type:    "error",
				Message: "Пользователь не найден",
			},
		},
		{
			name:           "некорректный ID",
			userID:         "abc",
			mockSetup:      func(m *MockUserService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Type:    "error",
				Message: "Ошибка преобразования id",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockUserService)
			tt.mockSetup(mockUserService)

			h := &Handler{
				services: &service.Service{
					User: mockUserService,
				},
			}

			req := httptest.NewRequest(http.MethodDelete, "/api/user/"+tt.userID, nil)

			// Создаем контекст с параметром chi
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.userID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()

			h.deleteUser(w, req)

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
			mockUserService.AssertExpectations(t)
		})
	}
}

func TestHandler_GetAllSportCategory(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "успешное получение спортивных категорий",
			mockSetup: func(m *MockUserService) {
				m.On("GetAllSportCategory", mock.Anything).Return([]domain.SportCategory{
					{ID: 1, Title: "Новичок"},
					{ID: 2, Title: "3 разряд"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.SportCategory{
				{ID: 1, Title: "Новичок"},
				{ID: 2, Title: "3 разряд"},
			},
		},
		{
			name: "ошибка получения категорий",
			mockSetup: func(m *MockUserService) {
				m.On("GetAllSportCategory", mock.Anything).Return([]domain.SportCategory{}, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения спортивных категорий",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockUserService)
			tt.mockSetup(mockUserService)

			h := &Handler{
				services: &service.Service{
					User: mockUserService,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/api/user/categories", nil)
			w := httptest.NewRecorder()

			h.getAllSportCategory(w, req)

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
			mockUserService.AssertExpectations(t)
		})
	}
}

func TestHandler_CheckToken(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:  "валидный токен",
			token: "Bearer valid-token-123",
			mockSetup: func(m *MockUserService) {
				m.On("CheckToken", mock.Anything, "valid-token-123").Return(true, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"valid": true,
			},
		},
		{
			name:  "невалидный токен",
			token: "Bearer invalid-token-123",
			mockSetup: func(m *MockUserService) {
				m.On("CheckToken", mock.Anything, "invalid-token-123").Return(false, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"valid": false,
			},
		},
		{
			name:  "ошибка проверки токена",
			token: "Bearer token-123",
			mockSetup: func(m *MockUserService) {
				m.On("CheckToken", mock.Anything, "token-123").Return(false, errors.New("ошибка сервера"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения токена",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockUserService)
			tt.mockSetup(mockUserService)

			h := &Handler{
				services: &service.Service{
					User: mockUserService,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/api/user/auth", nil)
			req.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()

			h.checkToken(w, req)

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
			mockUserService.AssertExpectations(t)
		})
	}
}

func TestHandler_GetAllAlpinists(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "успешное получение альпинистов",
			mockSetup: func(m *MockUserService) {
				m.On("GetAllAlpinists", mock.Anything).Return([]domain.User{
					{ID: 1, Name: "Иван Петров", Username: "ivan"},
					{ID: 2, Name: "Анна Сидорова", Username: "anna"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.User{
				{ID: 1, Name: "Иван Петров", Username: "ivan"},
				{ID: 2, Name: "Анна Сидорова", Username: "anna"},
			},
		},
		{
			name: "ошибка получения альпинистов",
			mockSetup: func(m *MockUserService) {
				m.On("GetAllAlpinists", mock.Anything).Return([]domain.User{}, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения альпинистов",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := new(MockUserService)
			tt.mockSetup(mockUserService)

			h := &Handler{
				services: &service.Service{
					User: mockUserService,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/api/alpinists", nil)
			w := httptest.NewRecorder()

			h.getAllAlpinists(w, req)

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
			mockUserService.AssertExpectations(t)
		})
	}
}
