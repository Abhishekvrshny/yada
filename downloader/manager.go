package downloader

import (
	"errors"
	"github.com/Abhishekvrshny/yada/models"
)

type Manager struct {
	downloads map[string]Downloader
}

func NewManager() *Manager{
	return &Manager{make(map[string]Downloader)}
}

func (mgr *Manager) NewDownload(request models.DownloadRequest) (Downloader, error) {
	if request.Type == "serial" {
		s, err := NewSerial(request.Files)
		mgr.downloads[s.GetID()] = s
		return s, err
	} else {
		return nil, errors.New("unknown type of download")
	}
}

func (mgr *Manager) GetStatus(id string) *models.StatusResponse {
	downloader, ok := mgr.downloads[id]
	if !ok {
		return nil
	} else {
		return &models.StatusResponse{
			ID: downloader.GetID(),
			StartTime:downloader.GetStartTime(),
			EndTime:downloader.GetEndTime(),
			Status:downloader.GetStatus(),
		}
	}
}