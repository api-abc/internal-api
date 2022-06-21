package adj

import (
	"context"
	"database/sql"
	"fmt"

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
	fmt.Println("ADJ - Process Query")
	rows, err := repo.database.Query(query, name)
	if err != nil {
		fmt.Println("ADJ REPO:", err)
		return dats
	}
	fmt.Println("ADJ - Process Query Done")
	// helper.HandlePanic(err)
	defer rows.Close()

	fmt.Println("ADJ - Process Rows")
	for rows.Next() {
		var data domain.Data
		err := rows.Scan(&data.Name, &data.Age, &data.Status, &data.JobDetails, &data.WorkerUpdate)
		if err != nil {
			fmt.Println("ADJ SCAN:", err)
			return dats
		}
		// helper.HandlePanic(err)
		dats = append(dats, &data)
	}
	return dats
}
