package insert

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/model/domain"
)

type IDataInsert interface {
	Insert(ctx context.Context, data domain.Data) error
	GetInserted(ctx context.Context, tx *sql.Tx) int
}
