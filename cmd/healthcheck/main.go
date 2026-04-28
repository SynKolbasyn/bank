package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/SynKolbasyn/bank/config"
)

func main() {
	config, err := config.LoadServer()
	if err != nil {
		panic(err)
	}

	url := url.URL{
		Scheme: "http",
		Host: config.Address(),
		Path: "health",
	}
    r, err := http.Get(url.String())
    if (err != nil) || (r.StatusCode != http.StatusOK) {
		os.Exit(1)
	}
}
