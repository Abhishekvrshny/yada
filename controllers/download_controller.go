package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Abhishekvrshny/yada/downloader"
	"github.com/Abhishekvrshny/yada/models"
	"github.com/Abhishekvrshny/yada/utils"
	"github.com/Abhishekvrshny/yada/yadaerror"
)

// DownloadController handles various download APIs
type DownloadController struct {
	DownloadMgr *downloader.Manager
}

// Download is the HTTP handler for download request
func (dc *DownloadController) Download(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		yErr := yadaerror.New(err.Error(), yadaerror.HTTP_READ_BODY_FAILED)
		setError(w, yErr)
		return
	}
	req := models.DownloadRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		yErr := yadaerror.New(err.Error(), yadaerror.JSON_UNMARSHAL_FAILED)
		setError(w, yErr)
		return
	}
	dMgr, err := dc.DownloadMgr.NewDownload(req)
	if err != nil {
		setError(w, err)
		return
	}
	dMgr.Download()
	res := &models.DownloadResponse{dMgr.GetID()}
	b, err := json.Marshal(res)
	if err != nil {
		yErr := yadaerror.New(err.Error(), yadaerror.JSON_MARSHAL_FAILED)
		setError(w, yErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// Status is the HTTP handler for status request
func (dc *DownloadController) Status(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromURI(r.RequestURI)
	if err != nil {
		setError(w, err)
		return
	}
	status, err := dc.DownloadMgr.GetStatus(id)
	if err != nil {
		setError(w, err)
		return
	}
	b, err := json.Marshal(status)
	if err != nil {
		yErr := yadaerror.New(err.Error(), yadaerror.JSON_MARSHAL_FAILED)
		setError(w, yErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// setError sets appropriate http return code.
// Handling only 400, 404 and 500 for now
func setError(w http.ResponseWriter, err error) {
	yErr := err.(yadaerror.Error)
	w.Header().Set("Content-Type", "application/json")
	if yErr.InternalCode >= 4000 && yErr.InternalCode < 4010 {
		w.WriteHeader(yadaerror.BAD_REQUEST)
	}
	if yErr.InternalCode >= 4040 && yErr.InternalCode < 4050 {
		w.WriteHeader(yadaerror.BAD_ROUTE)
	}
	if yErr.InternalCode >= 5000 && yErr.InternalCode < 5999 {
		w.WriteHeader(yadaerror.INTERNAL_ERROR)
	} else {
		w.WriteHeader(yadaerror.INTERNAL_ERROR)
	}
	w.Write(yErr.ToJSONBytes())
}
