package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/didsqq/crud-service-alpinizm/internal/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

func (h *Handler) getAllClimbs(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	query := req.URL.Query()
	mountainIDStr := query.Get("mountainId")
	categoryIDStr := query.Get("categoryId")

	var mountainID, categoryID int
	var err error

	if mountainIDStr != "" {
		mountainID, err = strconv.Atoi(mountainIDStr)
		if err != nil {
			h.respondError(w, http.StatusBadRequest, "Неверный mountainId", err)
			return
		}
	}

	if categoryIDStr != "" {
		categoryID, err = strconv.Atoi(categoryIDStr)
		if err != nil {
			h.respondError(w, http.StatusBadRequest, "Неверный categoryId", err)
			return
		}
	}
	log.Println("Get all climbs for", mountainID, categoryID)
	climbs, err := h.services.Climbs.GetAll(ctx, mountainID, categoryID)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения восхождений", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &climbs)
}

func (h *Handler) getClimb(w http.ResponseWriter, req *http.Request) {
	climbIDStr := chi.URLParam(req, "id")

	climbID, err := strconv.Atoi(climbIDStr)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Неверный climbId", err)
		return
	}

	ctx := req.Context()

	climb, err := h.services.Climbs.GetById(ctx, int64(climbID))
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения восхождения", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &climb)
}

func (h *Handler) recordAlpinistClimb(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	// log.Printf("Authorization header: %s", req.Header.Get("Authorization"))

	climbIDStr := chi.URLParam(req, "id")

	climbID, err := strconv.Atoi(climbIDStr)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Неверный climbId", err)
		return
	}

	// log.Printf("climbID: %d", climbID)

	_, claims, _ := jwtauth.FromContext(req.Context())

	// log.Printf("Claims from token: %+v", claims)

	id, ok := claims["id"].(float64)
	if !ok {
		h.respondError(w, http.StatusUnauthorized, "Некорректный токен: отсутствует id", nil)
		return
	}
	alpinistID := int64(id)

	ctx := req.Context()

	err = h.services.Climbs.RecordAlpinistClimb(ctx, alpinistID, int64(climbID))
	if err != nil {
		if errors.Is(err, repository.ErrAlpinistHasRegistered) {
			h.respondError(w, http.StatusBadRequest, "Вы уже зарегистрированы на это восхождение", err)
			return
		}
		h.respondError(w, http.StatusInternalServerError, "Ошибка записи восхождения", err)
		return
	}

	h.writeJSON(w, http.StatusOK, "Восхождение записано")
}

func (h *Handler) getAlpinistClimb(w http.ResponseWriter, req *http.Request) {

	_, claims, _ := jwtauth.FromContext(req.Context())

	id, ok := claims["id"].(float64)
	if !ok {
		h.respondError(w, http.StatusUnauthorized, "Некорректный токен: отсутствует id", nil)
		return
	}
	alpinistID := int64(id)

	ctx := req.Context()

	climbs, err := h.services.Climbs.GetAlpinistClimb(ctx, alpinistID)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения восхождений", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &climbs)
}

func (h *Handler) getAllCategoryOfDifficulty(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	categories, err := h.services.Climbs.GetAllCategoryOfDifficulty(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения категорий сложности", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &categories)
}

func (h *Handler) cancelAlpinistClimb(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	climbID := chi.URLParam(req, "id")

	_, claims, _ := jwtauth.FromContext(req.Context())

	id, ok := claims["id"].(float64)
	if !ok {
		h.respondError(w, http.StatusUnauthorized, "Некорректный токен: отсутствует id", nil)
		return
	}

	climbIDInt, err := strconv.Atoi(climbID)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Ошибка преобразования climb_id", err)
		return
	}

	if err := h.services.User.CancelAlpinistClimb(ctx, int64(id), int64(climbIDInt)); err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка отмены бронирования", err)
		return
	}

	h.respondSuccess(w, "Бронь отменена")
}
