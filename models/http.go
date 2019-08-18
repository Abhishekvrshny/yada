package models

type DownloadRequest struct {
	Type string   `json:"type"`
	URLs []string `json:"urls"`
}

type DownloadResponse struct {
	ID string `json:"id"`
}

type StatusResponse struct {
	ID           string `json:"id"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
	DownloadType string `json:"download_type"`
}
