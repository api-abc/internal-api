package update

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/domain"
)

type DataUpdateRepo struct {
	database *sql.DB
}

func NewDataUpdateRepo(db *sql.DB) IDataUpdate {
	return &DataUpdateRepo{
		database: db,
	}
}

func (repo *DataUpdateRepo) Update(ctx context.Context, tx *sql.Tx, data domain.Data) error {
	query := "UPDATE data SET name=$1, age=$2, job_details=$3, worker_update=$4 WHERE name=$5 AND status = true"
	_, err := tx.ExecContext(ctx, query, data.Name, data.Age, data.JobDetails, data.WorkerUpdate, data.Name)
	helper.HandlePanic(err)
	return nil
}

func (repo *DataUpdateRepo) GetUpdated(ctx context.Context, tx *sql.Tx) int {
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE status = true" //belum tau kalau updated check darimana
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
