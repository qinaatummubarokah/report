package model

type WriteReportRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Report struct {
	WriteReportRequest
}
