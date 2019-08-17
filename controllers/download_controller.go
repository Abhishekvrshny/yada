package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Abhishekvrshny/yada/constants"
	"github.com/Abhishekvrshny/yada/downloader"
	"github.com/Abhishekvrshny/yada/models"
	"github.com/Abhishekvrshny/yada/utils"
	"github.com/Abhishekvrshny/yada/yadaerror"
	"io/ioutil"
	"net/http"
)

type APIController struct {
	DownloadMgr *downloader.Manager
}

func (ac *APIController) Download(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(constants.INTERNAL_ERROR)
		w.Write(yadaerror.NewHTTPError(yadaerror.HTTPError{err.Error()}))
		return
	}
	req := models.DownloadRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(constants.INTERNAL_ERROR)
		w.Write(yadaerror.NewHTTPError(yadaerror.HTTPError{err.Error()}))
		return
	}
	dMgr, err := ac.DownloadMgr.NewDownload(req)
	if err != nil {
		w.WriteHeader(constants.INTERNAL_ERROR)
		w.Write(yadaerror.NewHTTPError(yadaerror.HTTPError{err.Error()}))
		return
	}
	dMgr.Download()
	res := &models.DownloadResponse{dMgr.GetID()}
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(constants.INTERNAL_ERROR)
		w.Write(yadaerror.NewHTTPError(yadaerror.HTTPError{err.Error()}))
		return
	}
	w.Write(b)
}

func (ac *APIController) Status(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromURI(r.RequestURI)
	if err != nil {
		w.WriteHeader(constants.BAD_ROUTE)
		w.Write(yadaerror.NewHTTPError(yadaerror.HTTPError{err.Error()}))
		return
	}
	status := ac.DownloadMgr.GetStatus(id)
	fmt.Println(status)
	b, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(constants.INTERNAL_ERROR)
		w.Write(yadaerror.NewHTTPError(yadaerror.HTTPError{err.Error()}))
		return
	}
	w.Write(b)
}

