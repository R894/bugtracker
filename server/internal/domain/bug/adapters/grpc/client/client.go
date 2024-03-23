package client

import (
	"context"

	"github.com/R894/bugtracker/internal/domain/bug/ports"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetBugsByProjectID(client ports.BugRepositoryServiceClient, request *ports.GetBugsByProjectIDRequest) (*ports.GetBugsByProjectIDResponse, error) {
	return client.GetBugsByProjectID(context.Background(), request)
}

func AssignBugTo(client ports.BugRepositoryServiceClient, request *ports.AssignBugToRequest) (*ports.EmptyResponse, error) {
	return client.AssignBugTo(context.Background(), request)
}

func GetBugByID(client ports.BugRepositoryServiceClient, request *ports.GetBugByIDRequest) (*ports.BugResponse, error) {
	return client.GetBugByID(context.Background(), request)
}

func NewGrpcClient(addr string) ports.BugRepositoryServiceClient {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil
	}
	return ports.NewBugRepositoryServiceClient(conn)
}
