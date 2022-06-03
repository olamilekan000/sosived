sosivapp:
	protoc -I./receiver/proto \
	--go_opt=module=github.com/olamilekan000/sosivio-app/receiver \
	--go_out=./receiver/proto \
	--go-grpc_opt=module=github.com/olamilekan000/sosivio-app/receiver \
	--go-grpc_out=./receiver/proto \
	./receiver/proto/*.proto

build:
	go build -o ./bin/ github.com/olamilekan000/sosivio-app/sender && \
	go build -o ./bin/ github.com/olamilekan000/sosivio-app/auth-service && \
	go build -o ./bin/ ./receiver/server/.
	go build -o ./bin/ ./receiver/client/.