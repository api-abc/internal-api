package update

import (
	"context"
	"database/sql"
	"encoding/json"

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

func (repo *DataUpdateRepo) Update(ctx context.Context, data domain.Data) (domain.Data, error) {
	query := "UPDATE data SET age=$1, job_details=$2, worker_update=$3 WHERE name=$4 AND status = true"
	marsh, err := json.Marshal(&data.JobDetails)
	helper.HandlePanic(err)
	_, err = repo.database.ExecContext(ctx, query, data.Age, marsh, data.WorkerUpdate, data.Name)
	helper.HandlePanic(err)
	return domain.Data{
		Name:         data.Name,
		Age:          data.Age,
		Status:       true,
		JobDetails:   data.JobDetails,
		WorkerUpdate: data.WorkerUpdate,
	}, nil
}

func (repo *DataUpdateRepo) GetUpdated(ctx context.Context) int {
	var dats []*domain.Data
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE status = true" //belum tau kalau updated check darimana
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
	return len(dats)
}
