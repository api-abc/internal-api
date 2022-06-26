package insert

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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
	marsh, err := json.Marshal(&data.JobDetails)
	helper.HandlePanic(err)
	result, err := repo.database.ExecContext(ctx, query, data.Name, data.Age, data.Status, marsh, data.WorkerUpdate)
	helper.HandlePanic(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("Insert Row ", rowsAffected, data.Name)
	helper.HandlePanic(err)
	if rowsAffected > 0 {
		return nil
	}
	return nil
}

func (repo *DataInsertRepo) GetInserted(ctx context.Context) []*domain.Data {
	var dats []*domain.Data
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE status = true"
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
