package update

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
	"github.com/api-abc/internal-api/services/update"
	"github.com/go-chi/chi"
)

type DataUsecaseUpdate struct {
	service update.IServiceUpdate
}

var ctx context.Context = context.Background()

func NewDataUsecaseUpdate(service update.IServiceUpdate) IDataUsecaseUpdate {
	return &DataUsecaseUpdate{
		service: service,
	}
}

func (upd *DataUsecaseUpdate) HandleUpdate(writer http.ResponseWriter, req *http.Request) {
	var request request.UpdateRequest
	name := chi.URLParam(req, "name")

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		helper.WriteOutput(writer, 500, helper.WriteStatusError(err, response.StatusInternalServerError))
		return
	}

	if !helper.ValidateAge(request.Age) {
		helper.WriteOutput(writer, 400, response.BodyResponse{Status: response.StatusBadRequest, Message: "Invalid Request", Data: nil})
		return
	}

	resp, err := upd.service.Update(ctx, request, name)
	if err != nil {
		switch err.Error() {
		case "data not found":
			helper.WriteOutput(writer, 404, helper.WriteStatusError(err, response.StatusNotFound))
		default:
			helper.WriteOutput(writer, 400, helper.WriteStatusError(err, response.StatusBadRequest))
		}
	} else {
		helper.WriteOutput(writer, 200, resp)
	}
}

func (upd *DataUsecaseUpdate) HandleGetUpdated(writer http.ResponseWriter, req *http.Request) {
	resp := upd.service.GetUpdate(ctx)
	helper.WriteOutputGet(writer, 200, resp)
}
