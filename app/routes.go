package app

import (
	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/usecase/delete"
	"github.com/api-abc/internal-api/usecase/insert"
	"github.com/api-abc/internal-api/usecase/update"

	"github.com/go-chi/chi"
)

func RoutesInsert(uc insert.IDataUsecaseInsert) *chi.Mux {
	r := chi.NewRouter()

	r.Use(helper.PanicRecovery)
	r.Route(`/internalapi/data`, func(r chi.Router) {
		r.Get(`/`, uc.HandleGetInserted)
		r.Post(`/`, uc.HandleInsert)
	})
	return r
}

func RoutesDelete(uc delete.IDataUsecaseDelete) *chi.Mux {
	r := chi.NewRouter()

	r.Use(helper.PanicRecovery)
	r.Route(`/internalapi/data`, func(r chi.Router) {
		r.Get(`/`, uc.HandleGetDeleted)
		r.Delete(`/{name}`, uc.HandleDelete)
	})

	return r
}

func RoutesUpdate(uc update.IDataUsecaseUpdate) *chi.Mux {
	r := chi.NewRouter()

	r.Use(helper.PanicRecovery)
	r.Route(`/internalapi/data`, func(r chi.Router) {
		r.Get(`/`, uc.HandleGetUpdated)
		r.Put(`/{name}`, uc.HandleUpdate)
	})

	return r
}
