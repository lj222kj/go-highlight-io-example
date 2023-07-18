package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/highlight/highlight/sdk/highlight-go"
	"log"
	"net/http"
)

func main() {
	var projectId string
	flag.StringVar(&projectId, "project", "", "the project id specified in your highlight.io workspace")
	flag.Parse()

	mux := http.NewServeMux()
	highlight.SetProjectID("2d1j6jdr")
	highlight.Start()
	defer highlight.Stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", 8080),
		Handler: mux,
	}

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		highlight.RecordError(r.Context(), errors.New("something went wrong"))
	})
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}
