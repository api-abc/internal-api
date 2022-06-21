package delete

import (
	"context"
	"net/http"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/services/delete"
	"github.com/go-chi/chi"
)

type DataUsecaseDelete struct {
	service delete.IServiceDelete
}

var ctx context.Context = context.Background()

func NewDataUsecaseDelete(service delete.IServiceDelete) IDataUsecaseDelete {
	return &DataUsecaseDelete{
		service: service,
	}
}

func (del *DataUsecaseDelete) HandleDelete(writer http.ResponseWriter, req *http.Request) {
	name := chi.URLParam(req, "name")
	resp := del.service.Delete(ctx, name)
	helper.WriteOutput(writer, 200, resp)
}

func (del *DataUsecaseDelete) HandleGetDeleted(writer http.ResponseWriter, req *http.Request) {
	resp := del.service.GetDelete(ctx)
	helper.WriteOutputGet(writer, 200, resp)
}
