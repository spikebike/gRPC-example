package main

import (
    "context"
    "fmt"
    "log"
    "net"

	"github.com/spikebike/gRPC-example/matrix"
	"google.golang.org/grpc"
)

type dataExchangeServer struct {
    matrix.UnimplementedDataExchangeServer 
}

func (s *dataExchangeServer) SendMatrix(ctx context.Context, matrix *matrix.Matrix) (*matrix.Status, error) {
    // Process received [][]byte (matrix.Rows)
    fmt.Println("Received matrix data:", matrix.Rows)

    // ... (Handle sending data to relevant peers based on your P2P logic)

    return &matrix.Status{Message: "Success"}, nil 
}

func main() {
    // ... Set up gRPC server (configure your listener, etc.)
    lis, err := net.Listen("tcp", ":8080") 
    // ... Error handling

    grpcServer := grpc.NewServer()
    matrix.RegisterDataExchangeServer(grpcServer, &dataExchangeServer{}) 
    grpcServer.Serve(lis)
}

