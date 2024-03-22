package client

import (
	"bugtracker/internal/domain/bug/ports"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetBugsByProjectID(c ports.BugRepositoryServiceClient, req *ports.GetBugsByProjectIDRequest) (*ports.GetBugsByProjectIDResponse, error) {
	return c.GetBugsByProjectID(context.TODO(), req)
}

func AssignBugTo(c ports.BugRepositoryServiceClient, req *ports.AssignBugToRequest) (*ports.EmptyResponse, error) {
	return c.AssignBugTo(context.TODO(), req)
}

func GetBugByID(c ports.BugRepositoryServiceClient, req *ports.GetBugByIDRequest) (*ports.BugResponse, error) {
	return c.GetBugByID(context.TODO(), req)
}

func GrpcClient(addr string) ports.BugRepositoryServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return ports.NewBugRepositoryServiceClient(conn)
}
