package main

import (
	metricsHandler "antonovxx/go-metrics/internal/handler"
	"antonovxx/go-metrics/internal/repository"
	"net/http"
)

func main() {
	storage := repository.CreateNewStorage()
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", metricsHandler.Update(storage))

	err := runServer(mux)

	if err != nil {
		panic(err)
	}
}

func runServer(mux *http.ServeMux) error {
	return http.ListenAndServe(":8080", mux)
}
