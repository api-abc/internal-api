package update

import (
	"context"
	"time"

	"github.com/api-abc/internal-api/model/domain"
	"github.com/api-abc/internal-api/model/request"
	"github.com/api-abc/internal-api/model/response"
	"github.com/api-abc/internal-api/repository/adj"
	"github.com/api-abc/internal-api/repository/update"
)

type ServiceUpdate struct {
	repo update.IDataUpdate
	adj  adj.IData
}

func NewUpdateService(adj adj.IData, repo update.IDataUpdate) IServiceUpdate {
	return &ServiceUpdate{
		repo: repo,
		adj:  adj,
	}
}

func (su *ServiceUpdate) Update(ctx context.Context, request request.UpdateRequest, name string) response.BodyResponse {
	//check if exist
	exist := su.adj.GetDataByName(ctx, name)
	if len(exist) == 0 {
		return response.BodyResponse{
			Status:  response.StatusNotFound,
			Message: "Data not found",
			Data:    nil,
		}
	}
	model := domain.Data{
		Name:         name,
		Age:          request.Age,
		JobDetails:   request.JobDetails,
		WorkerUpdate: time.Now(),
	}
	err := su.repo.Update(ctx, model)
	if err != nil {
		return response.BodyResponse{
			Status:  response.StatusInternalServerError,
			Message: "Failed to Update",
			Data:    nil,
		}
	}
	return response.BodyResponse{
		Status:  response.StatusOK,
		Message: "Update Success",
		Data:    nil,
	}
}

func (su *ServiceUpdate) GetUpdate(ctx context.Context) response.BodyResponseGet {
	count := su.repo.GetUpdated(ctx)
	return response.BodyResponseGet{
		Status:  response.StatusOK,
		Message: "Success fetch data",
		Data:    &count,
	}
}
