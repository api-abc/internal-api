package adj

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/model/domain"
)

type IData interface {
	GetDataByName(ctx context.Context, tx *sql.Tx, name string) domain.Data
}
