package server

import (
	"grpc-example/cmd"
	"grpc-example/conf"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGrpcServer() {
	//Listen to port by tcp protocol
	lis, err := net.Listen("tcp", conf.GetConfig().Vars.Port)
	if err != nil {
		log.Fatalln("Failed to listen ", err)
	}

	//Register service server
	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)
	usersserver := NewGServer()
	cmd.RegisterUserServiceServer(srv, usersserver)

	//Serve service
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
