HOME = C:\Users\cbott\Documents\chess-app
PROJECT_DIR = $(HOME)\servers\hive

PROTO = $(HOME)\proto\hive.proto

pb:
	protoc --go_out=. --go_opt=paths=$(PROJECT_DIR) --go-grpc_out=. --go-grpc_opt=paths=$(PROJECT_DIR) $(HOME)\proto\todo.proto
dev:
	cls && go run cmd\main.go