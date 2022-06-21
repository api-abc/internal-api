package update

import (
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

func (su *ServiceUpdate) Update(request request.UpdateRequest, name string) response.BodyResponse {
	res := response.BodyResponse{}
	return res
}

func (su *ServiceUpdate) GetUpdate() []response.BodyResponse {
	return nil
}
