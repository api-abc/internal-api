package update

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(request)
	if !helper.ValidateAge(request.Age) {
		helper.WriteOutput(writer, 400, response.BodyResponse{Status: response.StatusBadRequest, Message: "Invalid Request", Data: nil})
		return
	}
	resp := upd.service.Update(request, name)
	helper.WriteOutput(writer, 201, resp)
}

func (upde *DataUsecaseUpdate) HandleGetUpdated(writer http.ResponseWriter, req *http.Request) {

}
