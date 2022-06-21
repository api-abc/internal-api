package update

import (
	"context"

	"github.com/api-abc/internal-api/model/domain"
)

type IDataUpdate interface {
	Update(ctx context.Context, data domain.Data) (domain.Data, error)
	GetUpdated(ctx context.Context) int
}
