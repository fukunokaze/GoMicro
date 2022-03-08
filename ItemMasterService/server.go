package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fukunokaze/GoMicro/ItemMasterService/database"
	"github.com/fukunokaze/GoMicro/ItemMasterService/database/model"

	"github.com/fukunokaze/GoMicro/ItemMasterService/itemmasterpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	itemmasterpb.ItemMasterServiceServer
}

var (
	itemMasterRepository database.ItemMasterRepository = database.NewItemMasterRepository()
)

func main() {
	fmt.Println("Blog Service Started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	itemmasterpb.RegisterItemMasterServiceServer(s, &server{})
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

func (s *server) CreateItemMaster(ctx context.Context, in *itemmasterpb.CreateItemMasterRequest) (*itemmasterpb.CreateItemMasterResponse, error) {

	newItem := itemMasterRepository.CreateItemMaster(&model.ItemMaster{
		ItemName:   in.ItemMaster.ItemName,
		ItemNumber: in.ItemMaster.ItemNumber,
	})

	return &itemmasterpb.CreateItemMasterResponse{
		ItemMaster: &itemmasterpb.ItemMasterProto{
			ItemName:   newItem.ItemName,
			ItemNumber: newItem.ItemNumber,
		},
	}, nil
}
