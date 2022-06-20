package helper

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/api-abc/internal-api/model/response"
)

// untuk hadle panic
func HandlePanic(err error) {
	if err != nil {
		panic(err)
	}
}

// untuk handle commit dg database transaksional
func HandleCommit(tx *sql.Tx) {
	err := recover()
	if err != nil {
		e := tx.Rollback()
		HandlePanic(e)
		panic(err)
	} else {
		e := tx.Commit()
		HandlePanic(e)
	}
}

// untuk handle panic supaya aplikasi tidak mati
// dan keluar output internal server error
func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				resp := WriteStatusError(errors.New("internal server error"), response.StatusInternalServerError)
				WriteOutput(w, http.StatusInternalServerError, resp)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
