package models

// DownloadRequest : for download request payload
type DownloadRequest struct {
	Type string   `json:"type"`
	URLs []string `json:"urls"`
}

// DownloadResponse : for response of download request
type DownloadResponse struct {
	ID string `json:"id"`
}

// StatusResponse : for response of status request
type StatusResponse struct {
	ID           string            `json:"id"`
	StartTime    string            `json:"start_time"`
	EndTime      string            `json:"end_time"`
	Status       string            `json:"status"`
	DownloadType string            `json:"download_type"`
	Files        map[string]string `json:"files"`
}
