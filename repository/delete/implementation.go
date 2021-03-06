package delete

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

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
	result, err := repo.database.ExecContext(ctx, query, data.Status, data.Name)
	helper.HandlePanic(err)

	rowsAffected, err := result.RowsAffected()
	helper.HandlePanic(err)
	if rowsAffected > 0 {
		return nil
	}
	return errors.New("failed delete data")
}

func (repo *DataDeleteRepo) GetDeleted(ctx context.Context) []*domain.Data {
	var dats []*domain.Data
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE status = false"
	rows, err := repo.database.QueryContext(ctx, query)
	helper.HandlePanic(err)
	defer rows.Close()

	for rows.Next() {
		var data domain.Data
		var jobDetails []byte
		err := rows.Scan(&data.Name, &data.Age, &data.Status, &jobDetails, &data.WorkerUpdate)
		helper.HandlePanic(err)
		json.Unmarshal(jobDetails, &data.JobDetails)
		dats = append(dats, &data)
	}
	return dats
}
