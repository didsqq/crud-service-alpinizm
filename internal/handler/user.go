package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/handler/validate"
	"github.com/didsqq/crud-service-alpinizm/internal/repository"
	"github.com/go-chi/chi"
)

func (h *Handler) loginUser(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var user domain.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Ошибка парсинга json", err)
		return
	}

	ctx := req.Context()

	token, err := h.services.User.Login(ctx, user.Username, user.Password)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			h.respondError(w, http.StatusBadRequest, "Пользователь не найден", err)
		} else {
			h.respondError(w, http.StatusInternalServerError, "Ошибка входа пользователя", err)
		}
		return
	}

	h.respondSuccess(w, fmt.Sprintf("%s", token))
}

// @Summary      Создать пользователя
// @Tags         user
// @Description  Создаёт нового пользователя
// @Accept       json
// @Produce      json
// @Param        input body domain.UserInput true "Информация о пользователе"
// @Success      200 {string} string "Пользователь создан c id: {id}"
// @Failure      400 {string} string "Ошибка валидации или Username занят"
// @Failure      500 {string} string "Ошибка создания пользователя"
// @Router       /user [post]
func (h *Handler) createUser(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var user domain.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Ошибка парсинга json", err)
		return
	}

	if err := validate.ValidateUser(user); err != nil {
		h.respondError(w, http.StatusBadRequest, "Ошибка валидации", err)
		return
	}

	ctx := req.Context()

	id, err := h.services.User.Create(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNameExist) {
			h.respondError(w, http.StatusBadRequest, "Username занят", err)
		} else {
			h.respondError(w, http.StatusInternalServerError, "Ошибка создания пользователя", err)
		}
		return
	}

	h.respondSuccess(w, fmt.Sprintf("Пользователь создан c id: %d", id))
}

// @Summary      Получить пользователя по ID
// @Tags         user
// @Description  Возвращает пользователя по его ID
// @Accept       json
// @Produce      json
// @Param        id path int true "ID пользователя"
// @Success      200 {object} domain.User "Найденный пользователь"
// @Failure      400 {string} string "Ошибка преобразования id"
// @Failure      404 {string} string "Пользователь не найден"
// @Failure      500 {string} string "Ошибка получения пользователя"
// @Router       /user/{id} [get]
func (h *Handler) getUser(w http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Ошибка преобразования id", err)
		return
	}

	ctx := req.Context()

	user, err := h.services.User.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			h.respondError(w, http.StatusNotFound, "Пользователь не найден", err)
		} else {
			h.respondError(w, http.StatusInternalServerError, "Ошибка получения пользователя", err)
		}
		return
	}

	h.writeJSON(w, http.StatusOK, &user)
}

// @Summary      Удалить пользователя по ID
// @Tags         user
// @Description  Удаляет пользователя по его ID
// @Accept       json
// @Produce      json
// @Param        id path int true "ID пользователя"
// @Success      200 {string} string "Пользователь успешно удалён"
// @Failure      400 {string} string "Ошибка преобразования id"
// @Failure      404 {string} string "Пользователь не найден"
// @Failure      500 {string} string "Ошибка при удалении"
// @Router       /user/{id} [delete]
func (h *Handler) deleteUser(w http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Ошибка преобразования id", err)
		return
	}

	ctx := req.Context()

	if err := h.services.User.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			h.respondError(w, http.StatusNotFound, "Пользователь не найден", err)
		} else {
			h.respondError(w, http.StatusInternalServerError, "Ошибка при удалении", err)
		}
		return
	}

	h.respondSuccess(w, "Пользователь успешно удалён")
}

func (h *Handler) getAllSportCategory(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	c, err := h.services.User.GetAllSportCategory(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения спортивных категорий", err)
		return
	}

	h.writeJSON(w, http.StatusOK, &c)
}

func (h *Handler) checkToken(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	token := req.Header.Get("authorization")

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	b, err := h.services.User.CheckToken(ctx, token)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Ошибка получения токена", err)
		return
	}

	h.writeJSON(w, http.StatusOK, map[string]interface{}{
		"valid": b,
	})
}
