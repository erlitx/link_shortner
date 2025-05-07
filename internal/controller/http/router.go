package http

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi/v5"
	ver1 "github.com/erlitx/link_shortner/internal/controller/http/v1"
	"github.com/erlitx/link_shortner/internal/usecase"
	"net/http/pprof"
)


// Function to add pprof routes to Chi router
func RegisterPprofRoutes(r chi.Router) {
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
}


func ProfileRouter(r *chi.Mux, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	r.Handle("/metrics", promhttp.Handler())
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/create_shortlink", v1.CreateShortURL)
			RegisterPprofRoutes(r)
			
		})
	})

	r.Get("/{short_url}", v1.RedirectByShortURL)

}
