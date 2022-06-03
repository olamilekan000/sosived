package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/olamilekan000/sosivio-app/auth-service/helpers"
	"github.com/olamilekan000/sosivio-app/auth-service/interfaces"
	"github.com/olamilekan000/sosivio-app/auth-service/repository"

	pb "github.com/olamilekan000/sosivio-app/receiver/proto"
)

type MessagesService struct {
	serviceCaller   interfaces.ServiceCaller
	messagingClient pb.MessagingClient
	rpcMethods      interfaces.RpcServiceMethonds
}

func (m MessagesService) SendMessage(authToken string) (*helpers.SuccessResponse, *helpers.AppError) {

	fmt.Println("Verifying user")

	splitedToken := strings.Split(authToken, " ")

	if splitedToken[0] != "Bearer" {
		return nil, helpers.NewAuthorizationError("Inavlid Token")
	}

	resp, err := m.serviceCaller.MakeServiceRequest(splitedToken[1], "GET", "http://auth:8989/auth/verify", nil)

	if err != nil {
		log.Printf("HTTP Request to Micro service failed: %s", err)
		return nil, helpers.NewUnexpectedError("An unexpected error occured")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return nil, helpers.NewUnexpectedError("An unexpected error occured")
	}

	var response repository.APIResponse

	json.Unmarshal(bodyBytes, &response)

	if response.Code != http.StatusOK {
		return nil, helpers.NewAuthorizationError(response.Message)
	}

	messages := []*pb.MessageItem{
		{Message: "May the Force be with you."},
		{Message: "Hakuna Matata"},
		{Message: "Do. Or do not. There is no try."},
		{Message: "Feel the force!"},
		{Message: "Your path you must decide."},
	}

	go m.rpcMethods.ClientSendMessages(m.messagingClient, messages)

	return helpers.Success("Processing Messages", response.Data), nil
}

func NewMessageService(
	serviceCaller interfaces.ServiceCaller,
	messagingClient pb.MessagingClient,
	rpcMethods interfaces.RpcServiceMethonds,
) MessagesService {
	return MessagesService{
		serviceCaller,
		messagingClient,
		rpcMethods,
	}
}
