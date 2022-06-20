package insert

import "net/http"

type IDataUsecaseInsert interface {
	HandleInsert(writer http.ResponseWriter, req *http.Request)
	HandleGetInserted(writer http.ResponseWriter, req *http.Request)
}
