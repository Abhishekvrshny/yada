package constants

const (
	INTERNAL_ERROR = 500
	BAD_REQUEST = 400
	BAD_ROUTE = 404
)

type Status string

const (
	QUEUED Status = "QUEUED"
	FAILED Status = "FAILED"
	SUCCESSFUL Status = "SUCCESSFUL"
)

const DOWNLOADPATH = "/tmp/"
