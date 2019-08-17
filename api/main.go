package main

import (
	"github.com/Abhishekvrshny/yada/controllers"
	"github.com/Abhishekvrshny/yada/downloader"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	apiController := controllers.APIController{downloader.NewManager()}

	mux.HandleFunc("/health", controllers.Health)
	mux.HandleFunc("/download", apiController.Download)
	mux.HandleFunc("/status/", apiController.Status)

	http.ListenAndServe(":8081", mux)
}

