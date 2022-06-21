package delete

import (
	"net/http"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/services/delete"
	"github.com/go-chi/chi"
)

type DataUsecaseDelete struct {
	service delete.IServiceDelete
}

func NewDataUsecaseDelete(service delete.IServiceDelete) IDataUsecaseDelete {
	return &DataUsecaseDelete{
		service: service,
	}
}

func (del *DataUsecaseDelete) HandleDelete(writer http.ResponseWriter, req *http.Request) {
	name := chi.URLParam(req, "name")
	resp := del.service.Delete(name)
	helper.WriteOutput(writer, 201, resp)
}

func (del *DataUsecaseDelete) HandleGetDeleted(writer http.ResponseWriter, req *http.Request) {

}
