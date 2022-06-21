package delete

import (
	"github.com/api-abc/internal-api/model/response"
)

type IServiceDelete interface {
	Delete(name string) response.BodyResponse
	GetDelete() response.BodyResponse
}
