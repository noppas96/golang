package main

import (
	"context"
	"fmt"
	"github.com/noppas96/golang/generated"
	"log"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.DialContext(context.Background(), ":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dail: %v", err)
	}

	client := hello.NewHelloClient(clientConn)
	sayResp, err := client.Say(context.Background(), &hello.SayRequest{Name: "Por"})
	if err != nil {
		log.Fatalf("failed to say: %v", err)
	}
	fmt.Println(sayResp.Message)
}
