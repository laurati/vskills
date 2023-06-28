package main

import (
	"fmt"
	"net/http"
	"vskills/hydra/hlogger"
)

func main() {

	logger := hlogger.GetInstance()
	logger.Println("Starting hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)

}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()

	fmt.Fprint(w, "Welcome to the hydra software system")

	logger.Println("Received an http request on root url")
}
