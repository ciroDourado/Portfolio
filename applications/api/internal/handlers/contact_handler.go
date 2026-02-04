package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/cirodourado/portfolio/applications/api/internal/worker"
)

type ContactHandler struct {
	queueClient *asynq.Client
}

func NewContactHandler(queueClient *asynq.Client) *ContactHandler {
	return &ContactHandler{queueClient: queueClient}
}

func (ch *ContactHandler) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	var req worker.ContactPayload

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task, err := worker.NewContactTask(req.Name, req.Email, req.Message)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	if _, err := ch.queueClient.Enqueue(task, asynq.MaxRetry(5)); err != nil {
		http.Error(w, "Error queuing task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}