# Sosived

Testing gRPC

## Technologies

- Web/Server framework with [Golang](https://go.dev/).
- API Networking [gRPC](https://grpc.io/).
- API Networking [REST]
- Containerization [Docker](https://www.docker.com/)

## Getting started

```sh
# Clone the project
git clone https://github.com/olamilekan000/sosived.git
cd sosived

# Build and Run App
Docker-compose up or docker-compose build && docker-compose up

```

### Documentation

User Login

- API Endpoint

```sh
http://localhost:8989/auth/login

```

- Request Payload

```sh
  POST {
	"username": "admin",
	"password": "password"
}
```

- This should return a token
```sh 

{
	"data": {
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQyMTQxMzcsImlhdCI6MTY1NDIxMjMzNywiaXNzIjoiYWRtaW4ifQ.bQDGfyVZwAFS7TkjyZf1RDSMWWyV0IlzeFVl8AUMRog"
	},
	"statusCode": 200,
	"message": "Authentication was successful"
}
```

Send Message - This basically just triggers the gRPC client stream messages to the RPC server and in return to the streams, a single response is gotten. (Client Streaming)

- API Endpoint

```sh
http://localhost:8188/send-message

```

- Request Payload

```sh
  POST {}
```

- This should return a message and the logged in user name.
You an also see the data being streamed by the client and also being processed by the server.
```sh 

{
	"data": {
		"user": "admin"
	},
	"statusCode": 200,
	"message": "Processing Messages"
}
```

![alt text](message.processing.png)