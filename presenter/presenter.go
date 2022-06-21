package presenter

import (
	"database/sql"

	"github.com/api-abc/internal-api/repository/adj"
	insert_repo "github.com/api-abc/internal-api/repository/insert"
	insert_service "github.com/api-abc/internal-api/services/insert"
	insert_case "github.com/api-abc/internal-api/usecase/insert"

	update_repo "github.com/api-abc/internal-api/repository/update"
	update_service "github.com/api-abc/internal-api/services/update"
	update_case "github.com/api-abc/internal-api/usecase/update"

	delete_repo "github.com/api-abc/internal-api/repository/delete"
	delete_service "github.com/api-abc/internal-api/services/delete"
	delete_case "github.com/api-abc/internal-api/usecase/delete"
)

type PresenterMain struct {
	Db *sql.DB
}

var repo_adj adj.IData // kalau misal ada tambahan query lain misal GetDataByName dll (?)

func (presenter *PresenterMain) GetPresenterInsert() insert_case.IDataUsecaseInsert {
	repo_adj = adj.NewDataRepoAdj(presenter.Db)
	repo_insert := insert_repo.NewDataInsertRepo(presenter.Db)
	serv := insert_service.NewInsertService(repo_adj, repo_insert)
	return insert_case.NewDataUsecaseInsert(serv)
}

func (presenter *PresenterMain) GetPresenterDelete() delete_case.IDataUsecaseDelete {
	repo_adj = adj.NewDataRepoAdj(presenter.Db)
	repo_delete := delete_repo.NewDataDeleteRepo(presenter.Db)
	serv := delete_service.NewDeleteService(repo_adj, repo_delete)
	return delete_case.NewDataUsecaseDelete(serv)
}

func (presenter *PresenterMain) GetPresenterUpdate() update_case.IDataUsecaseUpdate {
	repo_adj = adj.NewDataRepoAdj(presenter.Db)
	repo_update := update_repo.NewDataUpdateRepo(presenter.Db)
	serv := update_service.NewUpdateService(repo_adj, repo_update)
	return update_case.NewDataUsecaseUpdate(serv)
}
