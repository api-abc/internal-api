package update

import (
	"context"

	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
)

type IServiceUpdate interface {
	Update(context.Context, request.UpdateRequest, string) response.BodyResponse
	GetUpdate(context.Context) response.BodyResponseGet
}
