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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMountainService - мок для сервиса гор
type MockMountainService struct {
	mock.Mock
}

func (m *MockMountainService) GetAll(ctx context.Context) ([]domain.Mountain, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Mountain), args.Error(1)
}

func TestHandler_GetAllMountains(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*MockMountainService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "успешное получение списка гор",
			mockSetup: func(m *MockMountainService) {
				m.On("GetAll", mock.Anything).Return([]domain.Mountain{
					{
						ID:            1,
						Title:         "Эльбрус",
						Height:        5642,
						MountainRange: "Кавказские горы",
					},
					{
						ID:            2,
						Title:         "Казбек",
						Height:        5047,
						MountainRange: "Кавказские горы",
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.Mountain{
				{
					ID:            1,
					Title:         "Эльбрус",
					Height:        5642,
					MountainRange: "Кавказские горы",
				},
				{
					ID:            2,
					Title:         "Казбек",
					Height:        5047,
					MountainRange: "Кавказские горы",
				},
			},
		},
		{
			name: "ошибка при получении списка гор",
			mockSetup: func(m *MockMountainService) {
				m.On("GetAll", mock.Anything).Return([]domain.Mountain{}, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения гор",
			},
		},
		{
			name: "пустой список гор",
			mockSetup: func(m *MockMountainService) {
				m.On("GetAll", mock.Anything).Return([]domain.Mountain{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []domain.Mountain{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок сервиса
			mockMountainService := new(MockMountainService)
			tt.mockSetup(mockMountainService)

			// Создаем хендлер с моком
			h := &Handler{
				services: &service.Service{
					Mountains: mockMountainService,
				},
			}

			// Создаем тестовый HTTP запрос
			req := httptest.NewRequest(http.MethodGet, "/api/mountain", nil)
			w := httptest.NewRecorder()

			// Вызываем хендлер
			h.getAllMountains(w, req)

			// Проверяем статус код
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Проверяем тело ответа
			var response interface{}
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			// Преобразуем ожидаемый ответ в JSON для сравнения
			expectedJSON, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)
			var expectedResponse interface{}
			err = json.Unmarshal(expectedJSON, &expectedResponse)
			assert.NoError(t, err)

			assert.Equal(t, expectedResponse, response)

			// Проверяем, что все ожидаемые вызовы мока были выполнены
			mockMountainService.AssertExpectations(t)
		})
	}
}
