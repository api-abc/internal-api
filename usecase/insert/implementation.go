package insert

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
	"github.com/api-abc/internal-api/services/insert"
)

type DataUsecaseInsert struct {
	service insert.IServiceInsert
}

var ctx context.Context = context.Background()

func NewDataUsecaseInsert(service insert.IServiceInsert) IDataUsecaseInsert {
	return &DataUsecaseInsert{
		service: service,
	}
}

func (ins *DataUsecaseInsert) HandleInsert(writer http.ResponseWriter, req *http.Request) {
	var request request.InsertRequest
	err := json.NewDecoder(req.Body).Decode(&request) //insert value from http body to request
	if err != nil {
		helper.WriteOutput(writer, 500, helper.WriteStatusError(err, response.StatusInternalServerError))
		return
	}

	if !helper.ValidateName(request.Name) || !helper.ValidateAge(request.Age) { //validate request
		helper.WriteOutput(writer, 400, response.BodyResponse{Status: response.StatusBadRequest, Message: "Invalid Request", Data: nil})
		return
	}

	resp, err := ins.service.Create(ctx, request) //run service by request
	if err != nil {
		helper.WriteOutput(writer, 400, helper.WriteStatusError(err, response.StatusBadRequest))
	} else {
		helper.WriteOutput(writer, 201, resp)
	}
}

func (ins *DataUsecaseInsert) HandleGetInserted(writer http.ResponseWriter, req *http.Request) {
	resp := ins.service.GetInsert(ctx)
	helper.WriteOutputGet(writer, 200, resp)
}
