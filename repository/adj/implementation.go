package adj

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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

func (repo *DataRepoAdj) GetDataByName(ctx context.Context, name string) []*domain.Data {
	var dats []*domain.Data
	query := "SELECT name, age, status, job_details, worker_update FROM data WHERE name = $1 AND status = true"
	rows, err := repo.database.Query(query, name)
	if err != nil {
		fmt.Println("ADJ REPO:", err)
		return dats
	}
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
