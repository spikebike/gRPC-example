package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/spikebike/gRPC-example/matrix"
)

type server struct {
	pb.UnimplementedByteTransferServiceServer
}

func (s *server) SendBytes(ctx context.Context, in *pb.ByteArrayRequest) (*pb.ByteArrayResponse, error) {
	log.Printf("Received request")
	// logic to handle and respond to the byte array
	res := &pb.ByteArrayResponse{Data: in.Data}
	return res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterByteTransferServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
