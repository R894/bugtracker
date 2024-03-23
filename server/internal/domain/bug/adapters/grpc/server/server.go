package grpc

import (
	"log"
	"net"

	"github.com/R894/bugtracker/internal/domain/bug/core/service"
	"github.com/R894/bugtracker/internal/domain/bug/ports"

	"google.golang.org/grpc"
)

type grpcServer struct {
	bugService service.BugService
	ports.UnimplementedBugRepositoryServiceServer
}

func NewGrpcServer(bugService service.BugService) *grpcServer {
	return &grpcServer{
		bugService: bugService,
	}
}

func (srv *grpcServer) StartServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(lis)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	ports.RegisterBugRepositoryServiceServer(s, srv)
	log.Printf("listening at %s\n", addr)

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
