package delete

import (
	"context"

	"github.com/api-abc/internal-api/model/response"
)

type IServiceDelete interface {
	Delete(context.Context, string) response.BodyResponse
	GetDelete(context.Context) response.BodyResponseGet
}
