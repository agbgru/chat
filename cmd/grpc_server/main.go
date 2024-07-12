package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/agbgru/chat/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServer
}

// CreateChat creates new chat for users from request.
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest,
) (*desc.CreateChatResponse, error) {
	_ = ctx

	log.Printf("CreateChat called with usernames: %v\n", req.Usernames)

	return &desc.CreateChatResponse{Id: 1}, nil
}

// DeleteChat deletes chat with id from request.
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest,
) (*desc.DeleteChatResponse, error) {
	_ = ctx

	log.Printf("DeleteChat called with id: %v\n", req.Id)

	return &desc.DeleteChatResponse{}, nil
}

// SendMessage sends message to chat.
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest,
) (*desc.SendMessageResponse, error) {
	_ = ctx

	log.Printf("SendMessage called from user %s with text %s and time %s",
		req.From, req.Text, req.Timestamp.AsTime().Format(time.DateTime))

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
