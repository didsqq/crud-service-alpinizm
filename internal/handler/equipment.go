package handler

import "net/http"

func (h *Handler) getAllEquipment(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	equipments, err := h.services.Equipments.GetAll(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения снаряжений", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &equipments)
}
