package helper

import (
	"encoding/json"
	"net/http"
)

// untuk validasi type data request body nya sesuai ngga
// jika ngga sesuai return error
func DecodeRequest(request *http.Request, result interface{}) error {
	err := json.NewDecoder(request.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}
