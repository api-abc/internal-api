package update

import "net/http"

type IDataUsecaseUpdate interface {
	HandleUpdate(writer http.ResponseWriter, req *http.Request)
	HandleGetUpdated(writer http.ResponseWriter, req *http.Request)
}
