package app

import (
	"net/http"

	"github.com/api-abc/internal-api/configuration"
	"github.com/api-abc/internal-api/helper"
	"github.com/api-abc/internal-api/presenter"
)

func Run(di *configuration.DI) {
	pm := presenter.PresenterMain{
		Db: NewDB(di),
	}

	usecase_insert := pm.GetPresenterInsert()
	usecase_delete := pm.GetPresenterDelete()
	usecase_update := pm.GetPresenterUpdate()

	port := di.GetConfig().Port
	go func() {
		err := http.ListenAndServe(port.Insert, RoutesInsert(usecase_insert))
		helper.HandlePanic(err)
	}()
	go func() {
		err := http.ListenAndServe(port.Delete, RoutesDelete(usecase_delete))
		helper.HandlePanic(err)
	}()
	err := http.ListenAndServe(port.Update, RoutesUpdate(usecase_update))
	helper.HandlePanic(err)
}
