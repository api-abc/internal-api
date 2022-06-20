package adj

import (
	"context"
	"database/sql"

	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/model/domain"
)

type DataRepoAdj struct{}

func NewDataRepoAdj() IData {
	return &DataRepoAdj{}
}

func (repo *DataRepoAdj) GetDataByName(ctx context.Context, tx *sql.Tx, name string) domain.Data {
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE name = $1 AND status = true"
	rows, err := tx.QueryContext(ctx, query, name)
	helper.HandlePanic(err)
	defer rows.Close()

	var data domain.Data
	if rows.Next() {
		err := rows.Scan(&data.Name, &data.Age, &data.Status, &data.JobDetails, &data.WorkerUpdate)
		helper.HandlePanic(err)
	}
	return data
}
