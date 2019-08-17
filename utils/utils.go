package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func GetIDFromURI(uri string) (string,error) {
	tokens := strings.Split(uri, "/")
	if len(tokens) != 3 {
		return "", errors.New("invalid status URL")
	} else {
		return tokens[2], nil
	}
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func CreateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return nil
}