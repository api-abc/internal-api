package insert

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/domain"
)

type DataInsertRepo struct {
	database *sql.DB
}

func NewDataInsertRepo(db *sql.DB) IDataInsert {
	return &DataInsertRepo{
		database: db,
	}
}

func (repo *DataInsertRepo) Insert(ctx context.Context, data domain.Data) error {
	query := "INSERT INTO data(name, age, status, job_details, worker_update) VALUES($1,$2,$3,$4,$5)"
	marsh, err := json.Marshal(data.JobDetails)
	if err != nil {
		return err
	}
	_, err = repo.database.ExecContext(ctx, query, data.Name, data.Age, data.Status, marsh, data.WorkerUpdate)
	helper.HandlePanic(err)
	return nil
}

func (repo *DataInsertRepo) GetInserted(ctx context.Context, tx *sql.Tx) int {
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE status = true"
	rows, err := tx.QueryContext(ctx, query)
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
