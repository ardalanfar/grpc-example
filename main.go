package main

import (
	"Service_grpc/client"
	"Service_grpc/server"

	"flag"
	"strings"
)

func main() {
	op := flag.String("op", "c", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		server.RunGrpcServer()
	case "c":
		client.RunGrpcClient()
	}
}
