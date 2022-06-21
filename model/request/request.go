package request

import "github.com/api-abc/internal-api/model/domain"

type InsertRequest struct {
	Name       string     `json:"name"`
	Age        int        `json:"age"`
	JobDetails domain.Job `json:"job_details"`
}

type UpdateRequest struct {
	Age        int        `json:"age"`
	JobDetails domain.Job `json:"job_details"`
}
