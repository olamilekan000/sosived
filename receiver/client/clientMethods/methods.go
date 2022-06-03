package clientmethoods

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/olamilekan000/sosivio-app/receiver/proto"
)

func (cd *RpcServiceMethonds) ClientSendMessages(c pb.MessagingClient, data []*pb.MessageItem) {

	stream, err := c.ProcessMessage(context.Background())

	if err != nil {
		log.Fatalf("Could not stream messages: %v", err)
	}

	for _, v := range data {
		stream.Send(v)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("An error occured after streaming data: %v", err)
	}

	fmt.Printf("All streamed Messages are:\n\n %v", res.Message)
}
