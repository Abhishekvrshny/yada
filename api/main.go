package main

import (
	"net/http"

	"github.com/Abhishekvrshny/yada/constants"

	"github.com/Abhishekvrshny/yada/controllers"
	"github.com/Abhishekvrshny/yada/downloader"
)

func main() {
	mux := http.NewServeMux()
	downloadController := controllers.DownloadController{downloader.NewManager()}

	mux.HandleFunc("/health", controllers.Health)
	mux.HandleFunc("/downloads", downloadController.Download)
	mux.HandleFunc("/downloads/", downloadController.Status)

	fs := http.FileServer(http.Dir(constants.DOWNLOAD_PATH))
	mux.Handle("/files/", http.StripPrefix("/files", fs))

	http.ListenAndServe(":8081", mux)
}
