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
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes(tokenAuth *jwtauth.JWTAuth) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://crud-service-alpinizm.onrender.com", "https://alteregoo.netlify.app"}, // фронтенд
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "Authorization"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour / time.Second),
	}))

	r.Use(middleware.LoggingMiddleware)

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			// Публичные маршруты
			r.Post("/registration", h.createUser)
			r.Post("/login", h.loginUser)
			r.Get("/categories", h.getAllSportCategory)

			// Защищённые маршруты
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(tokenAuth))
				r.Use(jwtauth.Authenticator(tokenAuth))

				r.Get("/", h.getUser)
				r.Delete("/{id}", h.deleteUser)

				r.Get("/auth", h.checkToken)
			})
		})

		r.Route("/climb", func(r chi.Router) {
			r.Get("/", h.getAllClimbs)
			r.Get("/{id}", h.getClimb)
			r.Get("/category", h.getAllCategoryOfDifficulty)

			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(tokenAuth))
				r.Use(jwtauth.Authenticator(tokenAuth))

				r.Get("/reservation", h.getAlpinistClimb)
				r.Post("/{id}/record", h.recordAlpinistClimb)
				r.Delete("/{id}/cancel", h.cancelAlpinistClimb)
			})
		})

		r.Route("/equipment", func(r chi.Router) {
			r.Get("/", h.getAllEquipment)

			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(tokenAuth))
				r.Use(jwtauth.Authenticator(tokenAuth))

				r.Put("/{id}", h.updateAlpinistEquipment)
				r.Delete("/{id}", h.deleteAlpinistEquipment)

				r.Post("/{id}/record", h.recordAlpinistEquipment)
				r.Get("/reservation", h.getAlpinistEquipment)
				r.Post("/{id}/delete", h.deleteAlpinistEquipment)
			})
		})

		r.Route("/mountain", func(r chi.Router) {
			r.Get("/", h.getAllMountains)
		})

		r.Get("/alpinists", h.getAllAlpinists)
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
