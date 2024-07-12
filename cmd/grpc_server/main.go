package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/agbgru/chat/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServer
}

// CreateChat ...
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest,
) (*desc.CreateChatResponse, error) {
	return &desc.CreateChatResponse{Id: 1}, nil
}

// DeleteChat ...
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest,
) (*desc.DeleteChatResponse, error) {
	return &desc.DeleteChatResponse{}, nil
}

// SendMessage ...
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest,
) (*desc.SendMessageResponse, error) {
	return &desc.SendMessageResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
