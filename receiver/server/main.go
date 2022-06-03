package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/olamilekan000/sosivio-app/receiver/proto"
)

const (
	port = ":50059"
)

type MessagingServer struct {
	pb.MessagingServer
}

func main() {
	conn, err := net.Listen("tcp", port)

	if err != nil {
		log.Println("Couldnt't create serevr " + err.Error())
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMessagingServer(grpcServer, &MessagingServer{})

	fmt.Println("Listening for requests")

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal("Could not connect to the grpc server")
	}

}

func (s *MessagingServer) ProcessMessage(stream pb.Messaging_ProcessMessageServer) error {

	fmt.Println("Processing streams")

	allMessages := ""

	for {
		data, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("Data streams done")
			return stream.SendAndClose(&pb.MessageItem{
				Message: allMessages,
			})
		}

		allMessages += fmt.Sprintf("%s\n", data.Message)

	}
}
