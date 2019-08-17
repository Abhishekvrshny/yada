package downloader

type Downloader interface {
	Download()
	GetID() string
	GetStartTime() string
	GetEndTime() string
	GetStatus() string
}

