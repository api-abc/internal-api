package delete

import (
	"context"

	"github.com/api-abc/internal-api/model/domain"
)

type IDataDelete interface {
	Delete(ctx context.Context, data domain.Data) error
	GetDeleted(ctx context.Context) []*domain.Data
}
