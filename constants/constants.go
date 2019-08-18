package constants

type Status string

const (
	QUEUED     Status = "QUEUED"
	FAILED     Status = "FAILED"
	SUCCESSFUL Status = "SUCCESSFUL"
)

const DOWNLOAD_PATH = "/tmp/"

const MAX_WORKERS_PER_DOWNLOAD = 10

const (
	SERIAL     = "SERIAL"
	CONCURRENT = "CONCURRENT"
)
