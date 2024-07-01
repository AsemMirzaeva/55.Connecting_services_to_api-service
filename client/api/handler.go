package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"task/client/proto/tasks"
	"task/client/qrcode"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

type TaskClient struct {
	Client tasks.TaskServiceClient
}

func NewTaskClient(cl tasks.TaskServiceClient) *TaskClient {
	return &TaskClient{Client: cl}
}

func (t *TaskClient) CreateTask(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	var httpReq tasks.TaskRequest
	protojson.Unmarshal(bytes, &httpReq)

	ctx, cancel := context.WithTimeout(r.Context(), 6*time.Second)
	defer cancel()

	resp, err := t.Client.CreateTask(ctx, &httpReq)
	if err != nil {
		log.Println("failed CreateTask method:", err)
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}

	fmt.Println(resp)

	code := qrcode.CreateQrcode(resp.TaskId)

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(201)
	w.Write(code)
}

func (t *TaskClient) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var req *tasks.Empty
	var ListTasks []*tasks.Task

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	stream, err := t.Client.ListTasks(ctx, req)
	if err != nil {
		log.Println("Failed ListTasks method:", err)
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}
	for {
		task, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, "Request Denied !", http.StatusInternalServerError)
			return
		}
		ListTasks = append(ListTasks, task)
		fmt.Println("-------Task------")
		fmt.Println(task)
	}

	fmt.Println("All Tasks Received Succesfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(ListTasks); err != nil {
		log.Println("Unable to send a response :", err)
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}
}

func (t *TaskClient) CancelTask(w http.ResponseWriter, r *http.Request) {
	var cancelReqHttp *tasks.CancelRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = protojson.Unmarshal(bytes, cancelReqHttp)
	if err != nil {
		http.Error(w, "Request Denied", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 6*time.Second)
	defer cancel()

	resp, err := t.Client.CancelTask(ctx, cancelReqHttp)
	if err != nil {
		log.Println("Failed CancelTask Method")
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}

	fmt.Println(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Unable to send a response :", err)
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}

}
