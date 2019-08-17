package downloader

import (
	"fmt"
	"github.com/Abhishekvrshny/yada/constants"
	"github.com/Abhishekvrshny/yada/utils"
	"time"
)

type Serial struct {
	id string
	startTime time.Time
	endTime time.Time
	files []string
	fileMap map[string]string
	status constants.Status
}

func NewSerial(files []string) (Downloader, error){
	uuid, err := utils.UUID()
	if err != nil {
		return nil, err
	}
	return &Serial{
		id : uuid,
		files: files,
		fileMap:make(map[string]string),
	}, nil
}

func (s *Serial) Download() {
	s.startTime = time.Now()
	defer func() {s.endTime = time.Now()}()

	s.status = constants.QUEUED
	for _, url := range s.files {
		fileID, err := utils.UUID()
		if err != nil {
			return
		}
		filePath := fmt.Sprintf("%s%s/%s",constants.DOWNLOADPATH, s.id, fileID)
		err = utils.CreateDir(constants.DOWNLOADPATH+s.id)
		if err != nil {
			return
		}
		err = utils.DownloadFile(filePath, url)
		if err != nil {
			return
		}
		s.fileMap[url] = fileID
	}
	s.status = constants.SUCCESSFUL
}

func (s *Serial) GetID() string {
	return s.id
}

func (s *Serial) GetStartTime() string {
	return s.startTime.String()
}

func (s *Serial) GetEndTime() string {
	return s.endTime.String()
}

func (s *Serial) GetStatus() string {
	return string(s.status)
}
