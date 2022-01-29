package main

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("artifacts-api", artifactsAPI)
}

func artifactsAPI(w http.ResponseWriter, r *http.Request) {}
