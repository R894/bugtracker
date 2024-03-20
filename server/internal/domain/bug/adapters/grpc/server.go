package grpc

import (
	"bugtracker/internal/database"
	"bugtracker/internal/domain/bug/adapters/repository"
	"bugtracker/internal/domain/bug/ports"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	d ports.BugRepository
	ports.UnimplementedBugRepositoryServiceServer
}

func GrpcServer() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}

	addr := args[0]
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

	db, err := database.NewPostgresTestDB()
	if err != nil {
		log.Fatalf("failed to connect to db: %v\n", err)
	}

	ports.RegisterBugRepositoryServiceServer(s, &server{d: repository.NewPostgresBugRepository(db.Db)})

	log.Printf("listening at %s\n", addr)

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
