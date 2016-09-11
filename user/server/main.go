package main

import (
	pb "grpc-go-trials/user/protos"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetUserDetails(ctx context.Context, in *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	log.Println(in.Id)
	return &pb.GetUserDetailsResponse{Message: "OK", Name: "Joe Doe"}, nil
}

func (s *server) SetUserDetails(ctx context.Context, in *pb.SetUserDetailsRequest) (*pb.SetUserDetailsResponse, error) {
	log.Println(in.Id, in.Name)
	return &pb.SetUserDetailsResponse{Message: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	s.Serve(lis)
}
