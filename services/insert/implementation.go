package insert

import (
	"context"
	"errors"
	"time"

	"github.com/api-abc/internal-api/model/domain"
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
	"github.com/api-abc/internal-api/repository/adj"
	"github.com/api-abc/internal-api/repository/insert"
)

type ServiceInsert struct {
	repo insert.IDataInsert
	adj  adj.IData
}

func NewInsertService(adj adj.IData, repo insert.IDataInsert) IServiceInsert {
	return &ServiceInsert{
		repo: repo,
		adj:  adj,
	}
}

func (si *ServiceInsert) Create(ctx context.Context, request request.InsertRequest) (response.BodyResponse, error) {
	//check one if exist
	exist := si.adj.GetDataByName(ctx, request.Name)
	if len(exist) != 0 {
		return response.BodyResponse{
			Status:  response.StatusBadRequest,
			Message: "Already Exist",
			Data:    nil,
		}, errors.New("already exist")
	}

	model := domain.Data{
		Name:         request.Name,
		Age:          request.Age,
		JobDetails:   request.JobDetails,
		Status:       true,
		WorkerUpdate: time.Now().Add(10 * time.Second),
	}

	err := si.repo.Insert(ctx, model)
	if err != nil {
		return response.BodyResponse{
			Status:  response.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}, errors.New("already exist")

	}
	return response.BodyResponse{
		Status:  response.StatusOK,
		Message: "Success",
		Data:    &model,
	}, nil
}

func (si *ServiceInsert) GetInsert(ctx context.Context) response.BodyResponseGet {
	count := si.repo.GetInserted(ctx)
	return response.BodyResponseGet{
		Status:  response.StatusOK,
		Message: "Success fetch data",
		Data:    count,
	}
}
