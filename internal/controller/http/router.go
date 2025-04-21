package http

import (
	"github.com/go-chi/chi/v5"
	ver1 "gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/controller/http/v1"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/usecase"
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

	r.Route("/mbelogortsev/my-app/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/profile", v1.CreateProfile)
			r.Put("/profile", v1.UpdateProfile)
			r.Get("/profile/{id}", v1.GetProfile)
			r.Delete("/profile/{id}", v1.DeleteProfile)
			r.Get("/profile_pg", v1.CreatePGProfile)
			r.Get("/wborders", v1.GetOrders)
			RegisterPprofRoutes(r)

		})
	})
}
