package main

import (
	"net/http"

	"storiesservice/pkg/logger"
)

func main() {
	defer gDatabase.Close()
	registerHandlers()
	logger.Println("Starting server on 0.0.0.0:80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
