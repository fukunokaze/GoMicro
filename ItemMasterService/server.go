package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fukunokaze/GoMicro/ItemMasterService/itemmasterpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	itemmasterpb.ItemMasterServiceServer
}

func main() {
	fmt.Println("Blog Service Started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	itemmasterpb.RegisterBlogServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

}

func (s *server) CreateItemMaster(ctx context.Context, in *CreateItemMasterRequest) (*CreateItemMasterResponse, error) {

	return &itemmasterpb.CreateItemMaster{
		ItemMaster: &itemmasterpb.ItemMasterProto{
			ItemId:     1,
			ItemNumber: "1",
		},
	}, nil
}
