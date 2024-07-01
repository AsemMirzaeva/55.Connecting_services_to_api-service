package api

import (
	"net/http"
	"task/client/clientdial"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	client := clientdial.DialGrpcClient()
	handler := NewTaskClient(client)

	mux.HandleFunc("POST /create", handler.CreateTask)
	mux.HandleFunc("GET /tasks", handler.GetAllTasks)
	mux.HandleFunc("DELETE  /cancel", handler.CancelTask)

	return mux
}
