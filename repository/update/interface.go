package update

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/model/domain"
)

type IDataUpdate interface {
	Update(ctx context.Context, tx *sql.Tx, data domain.Data) error
	GetUpdated(ctx context.Context, tx *sql.Tx) int
}
