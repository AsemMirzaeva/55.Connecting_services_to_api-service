package repo

import (
	"database/sql"
	"log"
	pb "task/proto/taskpb"
	"time"

	"github.com/google/uuid"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

func (t *TaskRepo) CreateTask(req *pb.TaskRequest) (*pb.TaskResponse, error) {
	time := time.Now().Format(time.ANSIC)
	id := uuid.New().String()

	query := `
		insert into tasks (id, task, start_at)
		values($1, $2, $3)
	`
	_, err := t.DB.Exec(query, id, req.TaskDescription, time)
	if err != nil {
		log.Println("Error while creating a new task:", err)
		return nil, err
	}
	return &pb.TaskResponse{
		TaskId: id,
		Status: "proccess",
	}, nil
}

func (t *TaskRepo) GetAllTasks(req *pb.Empty) ([]*pb.Task, error) {
	query := `
		select id, task, start_at from tasks
	`
	rows, err := t.DB.Query(query)
	if err != nil {
		log.Println("Error while reading tasks:", err)
		return nil, err
	}
	defer rows.Close()
	var listTasks []*pb.Task

	for rows.Next() {
		var task pb.Task
		err := rows.Scan(&task.Id, &task.TaskName, &task.StartedAt)
		if err != nil {
			log.Println("Error reading from db:", err)
			return nil, err
		}
		listTasks = append(listTasks, &task)
	}

	return listTasks, nil
}

func (t *TaskRepo) DeleteTaskFromDatabase(req *pb.CancelRequest) (*pb.CancelResponse, error) {
	query := `
		delete from tasks where id = $1
	`
	_, err := t.DB.Exec(query, req.TaskId)
	if err != nil {
		log.Println("Error deleting tasks:")
		return nil, err
	}
	return &pb.CancelResponse{Status: "Task Canceled Succesfully "}, nil
}
