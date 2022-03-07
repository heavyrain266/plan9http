package main

import (
	"fmt"
	"net/http"

	"github.com/HeavyRain266/plan9http/router"
)

func p9http(listenAddr string) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: router.NewRouter(),
	}

	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)

	return s.ListenAndServe()
}
