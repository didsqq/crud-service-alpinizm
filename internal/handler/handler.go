package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	"github.com/didsqq/crud-service-alpinizm/internal/handler/middleware"
	"github.com/didsqq/crud-service-alpinizm/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	services  *service.Service
	tokenAuth *jwtauth.JWTAuth
}

func NewHandler(services *service.Service, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{
		services:  services,
		tokenAuth: tokenAuth,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // фронтенд
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour / time.Second),
	}))

	r.Use(middleware.LoggingMiddleware)

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/{id}", h.getUser)
			r.Delete("/{id}", h.deleteUser)
			r.Post("/registration", h.createUser)
			r.Post("/login", h.loginUser)
		})

		r.Route("/climbs", func(r chi.Router) {
			r.Get("/", h.getAllClimbs)
			// r.Get("/{id}", h.getClimb)
		})

		r.Route("/equipment", func(r chi.Router) {
			r.Get("/", h.getAllEquipment)
		})

		r.Route("/mountain", func(r chi.Router) {
			r.Get("/", h.getAllMountains)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://localhost:8080/swagger/doc.json"),
	))

	return r
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Ошибка кодирования объекта", http.StatusInternalServerError)
		log.Printf("Ошибка при кодировании JSON-ответа: %v", err)
	}
}

func (h *Handler) respondError(w http.ResponseWriter, code int, msg string, err error) {
	log.Printf("%s: %v", msg, err)
	resp := domain.ApiResponse{
		Code:    code,
		Type:    "error",
		Message: msg,
	}
	h.writeJSON(w, code, resp)
}

func (h *Handler) respondSuccess(w http.ResponseWriter, msg string) {
	resp := domain.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: msg,
	}
	h.writeJSON(w, http.StatusOK, resp)
}
