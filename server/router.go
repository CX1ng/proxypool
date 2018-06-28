package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewProxyPoolRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	registerV1Router(router)
	return router
}

func registerV1Router(router chi.Router) {
	router.Route("/api", func(r chi.Router) {
		r.Get("/proxyip", getProxyIPWithLimit)
	})
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}
