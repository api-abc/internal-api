package insert

import (
	"context"

	"github.com/api-abc/internal-api/model/domain"
)

type IDataInsert interface {
	Insert(ctx context.Context, data domain.Data) error
	GetInserted(ctx context.Context) int
}
