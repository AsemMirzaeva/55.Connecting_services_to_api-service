package clientdial

import (
	"log"
	"task/client/proto/tasks"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcClient() tasks.TaskServiceClient {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client:", err)
	}

	return tasks.NewTaskServiceClient(conn)
}
