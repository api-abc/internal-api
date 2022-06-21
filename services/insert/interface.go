package insert

import (
	"context"

	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
)

type IServiceInsert interface {
	Create(context.Context, request.InsertRequest) (response.BodyResponse, error)
	GetInsert(context.Context) response.BodyResponseGet
}
