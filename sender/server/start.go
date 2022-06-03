package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olamilekan000/sosivio-app/auth-service/helpers"
	methods "github.com/olamilekan000/sosivio-app/receiver/client/clientMethods"
	pb "github.com/olamilekan000/sosivio-app/receiver/proto"
	"github.com/olamilekan000/sosivio-app/sender/handlers"
	"github.com/olamilekan000/sosivio-app/sender/services"
)

const (
	senderPort = ":8188"
	port       = "receiver-server:50059"
)

func Start() {

	fmt.Println("Starting Sender Service")

	router := mux.NewRouter()
	NotFoundHandler := http.HandlerFunc(handlers.NotFound)

	conn, err := ConnectToRPCServer()

	if err != nil {
		panic(err)
	}

	fmt.Println("Client conneted")
	defer conn.Close()

	rpcServerConnection := pb.NewMessagingClient(conn)
	meth := methods.RpcServiceMethonds{}

	service := services.NewMessageService(
		&helpers.ServiceCaller{},
		rpcServerConnection,
		&meth,
	)
	controllers := handlers.Sender{Service: service}

	router.NotFoundHandler = NotFoundHandler

	router.HandleFunc("/send-message", controllers.SendMessage).Methods("POST", "OPTIONS")
	router.HandleFunc("/say-hello", handlers.SayHello).Methods("GET")

	fmt.Printf("Sender Service Started at port %s\n", senderPort)

	http.ListenAndServe(senderPort, router)

}
