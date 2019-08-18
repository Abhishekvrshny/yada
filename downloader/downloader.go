package downloader

import (
	"time"

	"github.com/Abhishekvrshny/yada/constants"
)

type Downloader interface {
	Download() error
	GetID() string
	GetStartTime() string
	GetEndTime() string
	GetStatus() string
	GetType() string
	GetFiles() map[string]string
}

type BaseDownloader struct {
	id           string
	startTime    time.Time
	endTime      time.Time
	urls         []string
	fileMap      map[string]string
	status       constants.Status
	downloadType string
}

func (bd *BaseDownloader) GetID() string {
	return bd.id
}

func (bd *BaseDownloader) GetStartTime() string {
	return bd.startTime.String()
}

func (bd *BaseDownloader) GetEndTime() string {
	return bd.endTime.String()
}

func (bd *BaseDownloader) GetStatus() string {
	return string(bd.status)
}

func (bd *BaseDownloader) GetType() string {
	return string(bd.downloadType)
}

func (bd *BaseDownloader) GetFiles() map[string]string {
	return bd.fileMap
}
