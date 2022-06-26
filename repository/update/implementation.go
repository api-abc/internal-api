package update

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

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

func (repo *DataUpdateRepo) Update(ctx context.Context, data domain.Data) error {
	query := "UPDATE data SET age=$1, job_details=$2 WHERE name=$3 AND status = true"
	marsh, err := json.Marshal(&data.JobDetails)
	helper.HandlePanic(err)

	result, err := repo.database.ExecContext(ctx, query, data.Age, marsh, data.Name)
	helper.HandlePanic(err)

	rowsAffected, err := result.RowsAffected()
	fmt.Println("Update Row ", rowsAffected, data.Name)
	helper.HandlePanic(err)
	if rowsAffected > 0 {
		return nil
	}
	return errors.New("failed update data")
}

func (repo *DataUpdateRepo) GetUpdated(ctx context.Context) []*domain.Data {
	var dats []*domain.Data
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE worker_update between (CURRENT_TIMESTAMP - interval '1 second') and (CURRENT_TIMESTAMP + interval '1 second') LIMIT 10"
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
