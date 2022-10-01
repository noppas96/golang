package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/noppas96/golang/generated"
	"google.golang.org/grpc"
)

type helloServer struct{}

func (hs *helloServer) Say(ctx context.Context, in *hello.SayRequest) (*hello.SayResponse, error) {
	return &hello.SayResponse{Message: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &helloServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
