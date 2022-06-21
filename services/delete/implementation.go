package delete

import (
	"context"

	"github.com/api-abc/internal-api/model/domain"
	"github.com/api-abc/internal-api/model/response"
	"github.com/api-abc/internal-api/repository/adj"
	"github.com/api-abc/internal-api/repository/delete"
)

type ServiceDelete struct {
	repo delete.IDataDelete
	adj  adj.IData
}

func NewDeleteService(adj adj.IData, repo delete.IDataDelete) IServiceDelete {
	return &ServiceDelete{
		repo: repo,
		adj:  adj,
	}
}

func (sd *ServiceDelete) Delete(name string) response.BodyResponse {
	var ctx context.Context
	model := domain.Data{
		Name:   name,
		Status: false,
	}
	err := sd.repo.Delete(ctx, model)
	if err != nil {
		return response.BodyResponse{
			Status:  response.StatusBadRequest,
			Message: "Failed to Delete",
			Data:    nil,
		}
	}
	return response.BodyResponse{
		Status:  response.StatusOK,
		Message: "Delete Success",
		Data:    nil,
	}
}

func (sq *ServiceDelete) GetDelete() response.BodyResponse {
	res := response.BodyResponse{}
	return res
}
