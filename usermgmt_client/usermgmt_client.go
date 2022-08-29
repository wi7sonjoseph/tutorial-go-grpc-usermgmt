package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type userData struct {
	name         string
	mobile       int32
	organization string
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = []userData{
		{"Alice", 956506896, "google"},
		{"Tom", 954506896, "yahoo"},
	}
	for _, r := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: r.name, Mobile: r.mobile, Organisation: r.organization})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
		NAME: %s
		MOBILE: %d
		ID: %d 
		ORGANIZATION: %s`, r.GetName(), r.GetMobile(), r.GetId(), r.GetOrganisation())
	}
}
