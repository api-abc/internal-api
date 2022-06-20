package response

import "github.com/api-abc/internal-api/model/domain"

type BodyResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    domain.Data `json:"data,omitempty"`
}

const (
	StatusOK                  = 1
	StatusCreated             = 2
	StatusBadRequest          = 3
	StatusNotFound            = 4
	StatusInternalServerError = 5
)
