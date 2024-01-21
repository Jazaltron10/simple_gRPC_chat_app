package chat

import (
	"log"

	"golang.org/x/net/context"
)


type Server struct{
	UnimplementedChatServiceServer
}

func(s *Server)  SayHello(ctx context.Context, message *MessageRequest)(*MessageResponse, error){
	log.Printf("Received message body from the client: %s", message.Body)
	return &MessageResponse{Body:"Hello From the Server!"}, nil
}