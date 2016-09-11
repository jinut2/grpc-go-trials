package main

import (
	pb "grpc-go-trials/user/protos"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	r1, err := c.SetUserDetails(context.Background(), &pb.SetUserDetailsRequest{Id: "007", Name: "Joe Doe"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Welcome %s", r1.Message)

	r, err := c.GetUserDetails(context.Background(), &pb.GetUserDetailsRequest{Id: "007"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Welcome %s", r.Name)
}
