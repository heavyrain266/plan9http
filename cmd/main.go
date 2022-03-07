package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	var (
		host = flag.String("host", "localhost", "host http address")
		port = flag.String("port", "8080", "port number for http listener")
	)

	flag.Parse()

	adrr := net.JoinHostPort(*host, *port)

	if err := p9http(adrr); err != nil {
		log.Fatal(err)
	}
}
