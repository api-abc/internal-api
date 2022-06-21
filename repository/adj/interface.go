package adj

import (
	"context"

	"github.com/api-abc/internal-api/model/domain"
)

type IData interface {
	GetDataByName(ctx context.Context, name string) []domain.Data
}
