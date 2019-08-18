package main

import (
	"net/http"

	"github.com/Abhishekvrshny/yada/controllers"
	"github.com/Abhishekvrshny/yada/downloader"
)

func main() {
	mux := http.NewServeMux()
	apiController := controllers.APIController{downloader.NewManager()}

	mux.HandleFunc("/health", controllers.Health)
	mux.HandleFunc("/downloads", apiController.Download)
	mux.HandleFunc("/downloads/", apiController.Status)

	http.ListenAndServe(":8081", mux)
}
