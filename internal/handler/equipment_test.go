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

// MockEquipmentService - мок для сервиса снаряжения
type MockEquipmentService struct {
	mock.Mock
}

func (m *MockEquipmentService) GetAll(ctx context.Context) ([]domain.Equipment, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Equipment), args.Error(1)
}

func (m *MockEquipmentService) RecordAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	args := m.Called(ctx, alpinistID, equipmentID)
	return args.Error(0)
}

func (m *MockEquipmentService) GetAlpinistEquipment(ctx context.Context, alpinistID int64) ([]domain.AlpinistEquipment, error) {
	args := m.Called(ctx, alpinistID)
	return args.Get(0).([]domain.AlpinistEquipment), args.Error(1)
}

func (m *MockEquipmentService) UpdateAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64, equipment domain.AlpinistEquipment) error {
	args := m.Called(ctx, alpinistID, equipmentID, equipment)
	return args.Error(0)
}

func (m *MockEquipmentService) DeleteAlpinistEquipment(ctx context.Context, alpinistID int64, equipmentID int64) error {
	args := m.Called(ctx, alpinistID, equipmentID)
	return args.Error(0)
}

func (m *MockEquipmentService) GetAllEquipmentAdmin(ctx context.Context) ([]domain.AlpinistsEquipments, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.AlpinistsEquipments), args.Error(1)
}

// MockService - мок для основного сервиса
type MockService struct {
	Equipments *MockEquipmentService
}

func TestHandler_GetAllEquipment(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*MockEquipmentService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "успешное получение списка снаряжения",
			mockSetup: func(m *MockEquipmentService) {
				m.On("GetAll", mock.Anything).Return([]domain.Equipment{
					{
						ID:                1,
						Title:             "Веревка",
						QuantityAvailable: 10,
						ImageUrl:          "http://example.com/rope.jpg",
						Description:       "Основная веревка 50м",
					},
					{
						ID:                2,
						Title:             "Каска",
						QuantityAvailable: 15,
						ImageUrl:          "http://example.com/helmet.jpg",
						Description:       "Защитная каска",
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: []domain.Equipment{
				{
					ID:                1,
					Title:             "Веревка",
					QuantityAvailable: 10,
					ImageUrl:          "http://example.com/rope.jpg",
					Description:       "Основная веревка 50м",
				},
				{
					ID:                2,
					Title:             "Каска",
					QuantityAvailable: 15,
					ImageUrl:          "http://example.com/helmet.jpg",
					Description:       "Защитная каска",
				},
			},
		},
		{
			name: "ошибка при получении списка снаряжения",
			mockSetup: func(m *MockEquipmentService) {
				m.On("GetAll", mock.Anything).Return([]domain.Equipment{}, errors.New("ошибка базы данных"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: domain.ApiResponse{
				Code:    http.StatusInternalServerError,
				Type:    "error",
				Message: "Ошибка получения снаряжений",
			},
		},
		{
			name: "пустой список снаряжения",
			mockSetup: func(m *MockEquipmentService) {
				m.On("GetAll", mock.Anything).Return([]domain.Equipment{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []domain.Equipment{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок сервиса
			mockEquipmentService := new(MockEquipmentService)
			tt.mockSetup(mockEquipmentService)

			// Создаем хендлер с моком
			h := &Handler{
				services: &service.Service{
					Equipments: mockEquipmentService,
				},
			}

			// Создаем тестовый HTTP запрос
			req := httptest.NewRequest(http.MethodGet, "/api/equipment", nil)
			w := httptest.NewRecorder()

			// Вызываем хендлер
			h.getAllEquipment(w, req)

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
			mockEquipmentService.AssertExpectations(t)
		})
	}
}
