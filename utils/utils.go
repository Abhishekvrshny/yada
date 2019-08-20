package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Abhishekvrshny/yada/yadaerror"
)

// UUID generates a new UUID
func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", yadaerror.New(err.Error(), yadaerror.RANDOM_READ_FAILED)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

// GetIDFromURI parses uri to get ID, for status requests
func GetIDFromURI(uri string) (string, error) {
	tokens := strings.Split(uri, "/")
	if len(tokens) != 3 {
		return "", yadaerror.New("invalid route", yadaerror.INVALID_ROUTE)
	} else {
		return tokens[2], nil
	}
}

// DownloadFile downloads a url to filepath
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return yadaerror.New(err.Error(), yadaerror.HTTP_GET_FAILED)
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return yadaerror.New(err.Error(), yadaerror.FILE_CREATE_FAILED)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return yadaerror.New(err.Error(), yadaerror.FILE_COPY_FAILED)
	}
	return nil
}

// CreateDir to create a dir
func CreateDir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return yadaerror.New(err.Error(), yadaerror.DIR_CREATE_FAILED)
		}
	}
	if err != nil {
		return yadaerror.New(err.Error(), yadaerror.DIR_STAT_FAILED)
	}
	return nil
}
