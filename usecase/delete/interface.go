package delete

import "net/http"

type IDataUsecaseDelete interface {
	HandleDelete(writer http.ResponseWriter, req *http.Request)
	HandleGetDeleted(writer http.ResponseWriter, req *http.Request)
}
