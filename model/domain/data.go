package domain

import "time"

type Data struct {
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Status       bool      `json:"status"`
	JobDetails   Job       `json:"job_details"`
	WorkerUpdate time.Time `json:"worker_update"`
}

type Job struct {
	Position            string `json:"position"`
	YearsWorkExperience int    `json:"years_work_experience"`
	WorkStatus          string `json:"work_status"`
}
