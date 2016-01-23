package router

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"golang.org/x/net/context"
	"net/http"
)

func Router() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(func(h chi.Handler) chi.Handler {
		return chi.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			ctx = context.WithValue(ctx, "example", true)
			h.ServeHTTPC(ctx, w, r)
		})
	})

	r.Get("/", apiIndex)

	return r
}

func apiIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root"))
}
