package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":50059"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// messages := []*pb.MessageItem{
	// 	{Message: "May the Force be with you."},
	// 	{Message: "Hakuna Matata"},
	// 	{Message: "Do. Or do not. There is no try."},
	// 	{Message: "Feel the force!"},
	// 	{Message: "Your path you must decide."},
	// }

	if err != nil {
		log.Printf("Couldn't connect %v", err.Error())
	}

	defer conn.Close()

	fmt.Println("Client conneted")

	// c := pb.NewMessagingClient(conn)

	// meth := clientmethoods.RpcServiceMethonds{}
	// meth.ClientSendMessages(c, messages)
}
