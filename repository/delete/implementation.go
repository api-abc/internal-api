package delete

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/domain"
)

type DataDeleteRepo struct {
	database *sql.DB
}

func NewDataDeleteRepo(db *sql.DB) IDataDelete {
	return &DataDeleteRepo{
		database: db,
	}
}

func (repo *DataDeleteRepo) Delete(ctx context.Context, data domain.Data) error {
	query := "UPDATE data SET status=$1 WHERE name=$2 AND status = true"
	_, err := repo.database.ExecContext(ctx, query, data.Status, data.Name)
	helper.HandlePanic(err)
	return nil
}

func (repo *DataDeleteRepo) GetDeleted(ctx context.Context) int {
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE status = false"
	rows, err := repo.database.QueryContext(ctx, query)
	helper.HandlePanic(err)
	defer rows.Close()

	var dats []domain.Data
	for rows.Next() {
		var data domain.Data
		err := rows.Scan(&data.Name, &data.Age, &data.Status, &data.JobDetails, &data.WorkerUpdate)
		helper.HandlePanic(err)
		dats = append(dats, data)
	}
	return len(dats)
}
