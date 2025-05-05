package handler

import "net/http"

func (h *Handler) getAllClimbs(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	users, err := h.services.Climbs.GetAll(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения восхождений", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &users)
}
