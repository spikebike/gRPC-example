package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/spikebike/gRPC-example/matrix"
)

const (
	grpcTimeout = 10 * time.Second
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewByteTransferServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)
	defer cancel()

	data := [][]byte{{1, 2, 3}, {4, 5, 6}} // You would have to replace it with actual byte arrays you want to send
	r, err := c.SendBytes(ctx, &pb.ByteArrayRequest{Data: data})
	if err != nil {
		log.Fatalf("Could not send byte array: %v\n", err)
	}
	log.Printf("Byte array: %v\n", r.Data)
}
