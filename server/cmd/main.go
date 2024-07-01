package main

import (
	"log"
	"net"
	dbconn "task/postgres"
	"task/proto/taskpb"
	"task/server/repo"
	"task/server/service"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := dbconn.OpenDb("postgres", "postgres://postgres:1234@localhost:5432/demo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("Error while listenining:", err)
	}
	defer lis.Close()

	task := repo.NewTaskRepo(db)
	server := service.NewTaskServer(task)

	grpcServer := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	log.Println("Server is listening on port: 9090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error with server:", err)
	}

}
