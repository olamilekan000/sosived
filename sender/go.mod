module github.com/olamilekan000/sosivio-app/sender

go 1.18

replace github.com/olamilekan000/sosivio-app/auth-service v0.0.0 => ../auth-service

replace github.com/olamilekan000/sosivio-app/receiver v0.0.0 => ../receiver

require (
	github.com/gorilla/mux v1.8.0
	github.com/olamilekan000/sosivio-app/auth-service v0.0.0
	github.com/olamilekan000/sosivio-app/receiver v0.0.0
	github.com/rs/cors v1.8.2
	google.golang.org/grpc v1.47.0
)

require github.com/felixge/httpsnoop v1.0.1 // indirect

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/handlers v1.5.1
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
