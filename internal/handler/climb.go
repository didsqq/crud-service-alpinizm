package handler

import (
	"log"
	"net/http"
	"strconv"
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
