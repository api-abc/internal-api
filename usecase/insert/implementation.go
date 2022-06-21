package insert

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
	"github.com/api-abc/internal-api/services/insert"
)

type DataUsecaseInsert struct {
	service insert.IServiceInsert
}

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
	fmt.Println(request)
	if !helper.ValidateName(request.Name) || !helper.ValidateAge(request.Age) { //validate request
		helper.WriteOutput(writer, 400, response.BodyResponse{Status: response.StatusBadRequest, Message: "Invalid Request", Data: nil})
		return
	}
	resp := ins.service.Create(request) //run service by request
	helper.WriteOutput(writer, 201, resp)
}

func (ins *DataUsecaseInsert) HandleGetInserted(writer http.ResponseWriter, req *http.Request) {
	resp := ins.service.GetInsert()
	helper.WriteOutput(writer, response.StatusOK, resp)
}
