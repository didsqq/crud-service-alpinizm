package handler

import "net/http"

func (h *Handler) getAllMountains(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	mountains, err := h.services.Mountains.GetAll(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения гор", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &mountains)
}
