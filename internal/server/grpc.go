package server

import (
	"log"
	"net"

	"github.com/1024casts/snake/internal/conf"

	"google.golang.org/grpc"

	v1 "github.com/1024casts/snake/api/grpc/user/v1"
	"github.com/1024casts/snake/internal/service"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(svc *service.Service) *grpc.Server {
	lis, err := net.Listen("tcp", conf.Conf.Grpc.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	v1.RegisterUserServiceServer(grpcServer, service.UserSvc)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
	log.Printf("serve grpc server is success, port:%s", conf.Conf.Grpc.Addr)

	return grpcServer
}
