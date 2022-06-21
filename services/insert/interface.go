package insert

import (
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
)

type IServiceInsert interface {
	Create(request.InsertRequest) response.BodyResponse
	GetInsert() response.BodyResponse
}
