package constants

type Status string

// download states
const (
	QUEUED     Status = "QUEUED"
	FAILED     Status = "FAILED"
	SUCCESSFUL Status = "SUCCESSFUL"
)

// root location to download files
const DOWNLOAD_PATH = "/tmp/"

// max no of goroutines spawned per download request
const MAX_WORKERS_PER_DOWNLOAD = 10

// types of download requests
const (
	SERIAL     = "SERIAL"
	CONCURRENT = "CONCURRENT"
)
