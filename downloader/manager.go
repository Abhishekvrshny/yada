package downloader

import (
	"strings"

	"github.com/Abhishekvrshny/yada/constants"
	"github.com/Abhishekvrshny/yada/models"
	"github.com/Abhishekvrshny/yada/yadaerror"
)

type Manager struct {
	downloads map[string]Downloader
}

func NewManager() *Manager {
	return &Manager{make(map[string]Downloader)}
}

func (mgr *Manager) NewDownload(request models.DownloadRequest) (Downloader, error) {
	if strings.ToUpper(request.Type) == constants.SERIAL {
		s, err := NewSerial(request.URLs)
		if err != nil {
			return nil, err
		}
		mgr.downloads[s.GetID()] = s
		return s, nil
	}
	if strings.ToUpper(request.Type) == constants.CONCURRENT {
		s, err := NewConcurrent(request.URLs)
		if err != nil {
			return nil, err
		}
		mgr.downloads[s.GetID()] = s
		return s, nil
	} else {
		return nil, yadaerror.New("unknown type of download", yadaerror.UNKNOWN_DOWNLOAD_TYPE)
	}
}

func (mgr *Manager) GetStatus(id string) (*models.StatusResponse, error) {
	downloader, ok := mgr.downloads[id]
	if !ok {
		return nil, yadaerror.New("unknown download ID", yadaerror.UNKNOWN_DOWNLOAD_ID)
	} else {
		return &models.StatusResponse{
			ID:           downloader.GetID(),
			StartTime:    downloader.GetStartTime(),
			EndTime:      downloader.GetEndTime(),
			Status:       downloader.GetStatus(),
			DownloadType: downloader.GetType(),
		}, nil
	}
}
