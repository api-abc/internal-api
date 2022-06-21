package adj

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/domain"
)

type DataRepoAdj struct {
	database *sql.DB
}

func NewDataRepoAdj(db *sql.DB) IData {
	return &DataRepoAdj{
		database: db,
	}
}

func (repo *DataRepoAdj) GetDataByName(ctx context.Context, name string) []domain.Data {
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE name = $1 AND status = true"
	rows, err := repo.database.QueryContext(ctx, query, name)
	helper.HandlePanic(err)
	defer rows.Close()

	var dats []domain.Data
	for rows.Next() {
		var data domain.Data
		err := rows.Scan(&data.Name, &data.Age, &data.Status, &data.JobDetails, &data.WorkerUpdate)
		helper.HandlePanic(err)
		dats = append(dats, data)
	}
	return dats
}
