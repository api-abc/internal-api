package helper

import (
	"encoding/json"
	"net/http"

	"github.com/api-abc/internal-api/model/domain"
	"github.com/api-abc/internal-api/model/response"
)

// menuliskan output json di usecase handler
func WriteOutput(writer http.ResponseWriter, code int, resp response.BodyResponse) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	output, _ := json.Marshal(resp)
	writer.Write([]byte(output))
}

// menuliskan output json di usecase handler
func WriteOutputGet(writer http.ResponseWriter, code int, resp response.BodyResponseGet) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	output, _ := json.Marshal(resp)
	writer.Write([]byte(output))
}

// membuat response untuk kasus berhasil (status 200, 201, dll)
func WriteStatusSuccess(result *domain.Data, status int, msg string) response.BodyResponse {
	resp := response.BodyResponse{
		Status:  status,
		Message: msg,
		Data:    result,
	}
	return resp
}

// membuat response untuk kasus error (status 400, 404, 500, dll)
func WriteStatusError(err error, status int) response.BodyResponse {
	resp := response.BodyResponse{
		Status:  status,
		Message: err.Error(),
		Data:    nil,
	}
	return resp
}
