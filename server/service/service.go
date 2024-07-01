package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "task/proto/taskpb"
	"task/server/repo"
	"time"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	TaskRepo *repo.TaskRepo
}

func NewTaskServer(tr *repo.TaskRepo) *TaskServer {
	return &TaskServer{TaskRepo: tr}
}

func (s *TaskServer) CreateTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	log.Println("task received :", req.TaskDescription)

	resp, err := s.TaskRepo.CreateTask(req)
	if err != nil {
		return nil, fmt.Errorf("request denied ")
	}

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return resp, ctx.Err()
		case <-time.After(1 * time.Second):
			log.Println("task is ongoing")
		}
	}
	return &pb.TaskResponse{TaskId: resp.TaskId, Status: "success"}, nil
}

func (s *TaskServer) ListTasks(req *pb.Empty, stream pb.TaskService_ListTasksServer) error {
	listTask, err := s.TaskRepo.GetAllTasks(req)
	if err != nil {
		return fmt.Errorf("request denied")
	}

	for i := 0; i < len(listTask); i++ {

		err := stream.Send(listTask[i])
		if err != nil {
			return fmt.Errorf("unable to send a response")
		}
	}

	stream.Send(&pb.Task{Id: "All task sended succesfully"})
	return nil
}

func (t *TaskServer) CancelTask(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error) {

	resp, err := t.TaskRepo.DeleteTaskFromDatabase(req)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no task found with ID: %s", req.TaskId)
	}
	if err != nil {
		return nil, fmt.Errorf("request denied")
	}

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			log.Println("client denied the response:")
			return nil, fmt.Errorf("unable to cancel task, timeout exceed ")
		case <-time.After(1 * time.Second):
			log.Println("in proccess: cancelling task")
		}
	}
	fmt.Println("Canceled Successfully")
	return resp, nil
}
