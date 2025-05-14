package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

func (h *Handler) getAllEquipment(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	equipments, err := h.services.Equipments.GetAll(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения снаряжений", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &equipments)
}

func (h *Handler) recordAlpinistEquipment(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	equipmentIDStr := chi.URLParam(req, "id")

	equipmentID, err := strconv.Atoi(equipmentIDStr)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Неверный equipmentId", err)
		return
	}

	_, claims, _ := jwtauth.FromContext(req.Context())

	alpinistID, ok := claims["id"].(float64)
	if !ok {
		h.respondError(w, http.StatusUnauthorized, "Некорректный токен: отсутствует id", nil)
		return
	}

	if err := h.services.Equipments.RecordAlpinistEquipment(ctx, int64(alpinistID), int64(equipmentID)); err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка записи снаряжения", err)
		return
	}

	h.writeJSON(w, http.StatusOK, "Снаряжение записано")
}

func (h *Handler) getAlpinistEquipment(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	_, claims, _ := jwtauth.FromContext(req.Context())

	alpinistID, ok := claims["id"].(float64)
	if !ok {
		h.respondError(w, http.StatusUnauthorized, "Некорректный токен: отсутствует id", nil)
		return
	}

	equipments, err := h.services.Equipments.GetAlpinistEquipment(ctx, int64(alpinistID))
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения снаряжения", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &equipments)
}

func (h *Handler) deleteAlpinistEquipment(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	equipmentIDStr := chi.URLParam(req, "id")

	equipmentID, err := strconv.Atoi(equipmentIDStr)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Неверный equipmentId", err)
		return
	}

	_, claims, _ := jwtauth.FromContext(req.Context())

	alpinistID, ok := claims["id"].(float64)
	if !ok {
		h.respondError(w, http.StatusUnauthorized, "Некорректный токен: отсутствует id", nil)
		return
	}

	if err := h.services.User.DeleteAlpinistEquipment(ctx, int64(alpinistID), int64(equipmentID)); err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка удаления снаряжения", err)
		return
	}

	h.writeJSON(w, http.StatusOK, "Снаряжение удалено")
}
