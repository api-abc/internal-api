package update

import (
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
)

type IServiceUpdate interface {
	Update(request request.UpdateRequest, name string) response.BodyResponse
	GetUpdate() []response.BodyResponse
}
