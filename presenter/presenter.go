package presenter

import (
	"database/sql"
	"golang-api-internal/repository/delete_repo"
	"golang-api-internal/repository/insert_repo"
	"golang-api-internal/repository/update_repo"
	"golang-api-internal/service/delete_service"
	"golang-api-internal/service/insert_service"
	"golang-api-internal/service/update_service"
	"golang-api-internal/usecase/delete_usecase"
	"golang-api-internal/usecase/insert_usecase"
	"golang-api-internal/usecase/update_usecase"

	"github.com/api-abc/internal-api/repository/adj"
)

type PresenterMain struct {
	Db *sql.DB
}

var repo_adj = adj.NewDataRepoAdj() // kalau misal ada tambahan query lain misal GetDataByName dll (?)

func (presenter *PresenterMain) GetPresenterInsert() insert_usecase.IDataUsecaseInsert {
	repo_insert := insert_repo.NewDataRepoInsert()
	serv := insert_service.NewDataServiceInsert(repo_adj, repo_insert, presenter.Db)
	return insert_usecase.NewDataUsecaseInsert(serv)
}

func (presenter *PresenterMain) GetPresenterDelete() delete_usecase.IDataUsecaseDelete {
	repo_delete := delete_repo.NewDataRepoDelete()
	serv := delete_service.NewDataServiceDelete(repo_adj, repo_delete, presenter.Db)
	return delete_usecase.NewDataUsecaseDelete(serv)
}

func (presenter *PresenterMain) GetPresenterUpdate() update_usecase.IDataUsecaseUpdate {
	repo_update := update_repo.NewDataRepoUpdate()
	serv := update_service.NewDataServiceUpdate(repo_adj, repo_update, presenter.Db)
	return update_usecase.NewDataUsecaseUpdate(serv)
}
