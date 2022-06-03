package interfaces

import (
	pb "github.com/olamilekan000/sosivio-app/receiver/proto"
)

type RpcServiceMethonds interface {
	ClientSendMessages(c pb.MessagingClient, data []*pb.MessageItem)
}
