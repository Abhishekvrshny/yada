package yadaerror

// supported HTTP error codes
const (
	INTERNAL_ERROR = 500
	BAD_REQUEST    = 400
	BAD_ROUTE      = 404
)

// internal error codes that map to
// 4xx HTTP codes
const (
	UNKNOWN_DOWNLOAD_TYPE = 4001
	UNKNOWN_DOWNLOAD_ID   = 4002
	INVALID_ROUTE         = 4040
)

// internal error codes that map to
// 5xx HTTP codes
const (
	HTTP_READ_BODY_FAILED = 5001
	JSON_UNMARSHAL_FAILED = 5002
	JSON_MARSHAL_FAILED   = 5003
	RANDOM_READ_FAILED    = 5004
	HTTP_GET_FAILED       = 5005
	FILE_CREATE_FAILED    = 5005
	FILE_COPY_FAILED      = 5006
	DIR_CREATE_FAILED     = 5007
	DIR_STAT_FAILED       = 5008
)
