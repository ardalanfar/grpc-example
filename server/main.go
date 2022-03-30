package server

import (
	"Service_grpc/cmd"
	"Service_grpc/conf"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func RunGrpcServer() {
	grpclog.Println("Starting Server...")
	lis, err := net.Listen("tcp", conf.GetConfig().Vars.Port)
	if err != nil {
		log.Fatalln("Failed to listen ", err)
	}

	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)
	usersserver := NewGServer()
	cmd.RegisterUserServiceServer(srv, usersserver)

	err = srv.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
