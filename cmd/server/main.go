package main

import (
	"fmt"
	"net"

	"github.com/gislihr/userservice/proto"
	"github.com/gislihr/userservice/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.WithError(err).Fatal("failed to listen")
	}

	grpcServer := grpc.NewServer()

	s := service.New()
	proto.RegisterUserServiceServer(grpcServer, s)

	log.WithField("port", port).Info("Listening...")
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
