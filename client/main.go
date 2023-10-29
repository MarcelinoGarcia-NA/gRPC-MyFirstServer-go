package main

import (
	"Hello/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := pb.NewHelloClient(conn)

	req := &pb.HelloResquest{
		Name: "John Doe",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
