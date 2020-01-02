package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ddddddO/work/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {
	addr := ":50051"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cred, err := credentials.NewServerTLSFromFile(
		"../data/server.crt",
		"../data/private.key",
	)
	if err != nil {
		log.Fatal(err)
	}

	//s := grpc.NewServer()
	s := grpc.NewServer(grpc.Creds(cred))
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
