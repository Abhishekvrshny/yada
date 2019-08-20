package downloader

import (
	"fmt"
	"time"

	"github.com/Abhishekvrshny/yada/constants"
	"github.com/Abhishekvrshny/yada/utils"
)

// Serial implements Downloader interface
// and is composed of BaseDownloader
type Serial struct {
	BaseDownloader
}

// NewSerial returns an instance of Serial downloader
func NewSerial(urls []string) (Downloader, error) {
	uuid, err := utils.UUID()
	if err != nil {
		return nil, err
	}
	return &Serial{
		BaseDownloader{
			id:           uuid,
			urls:         urls,
			fileMap:      make(map[string]string),
			downloadType: constants.SERIAL,
		},
	}, nil
}

// Download the files serially
func (s *Serial) Download() error {
	s.startTime = time.Now()
	defer func() { s.endTime = time.Now() }()

	s.status = constants.QUEUED
	for _, url := range s.urls {
		fileID, err := utils.UUID()
		if err != nil {
			s.status = constants.FAILED
			return err
		}
		filePath := fmt.Sprintf("%s%s/%s", constants.DOWNLOAD_PATH, s.id, fileID)
		err = utils.CreateDir(constants.DOWNLOAD_PATH + s.id)
		if err != nil {
			s.status = constants.FAILED
			return err
		}
		err = utils.DownloadFile(filePath, url)
		if err != nil {
			s.status = constants.FAILED
			return err
		}
		s.fileMap[url] = filePath
	}
	s.status = constants.SUCCESSFUL
	return nil
}
