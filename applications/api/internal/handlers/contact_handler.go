package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/cirodourado/portfolio/applications/api/internal/worker"
)

type ContactHandler struct {
	queueClient *asynq.Client
	validate    *validator.Validate
}

func NewContactHandler(queueClient *asynq.Client) *ContactHandler {
	return &ContactHandler{
	    queueClient: queueClient,
	    validate:    validator.New(),
    }
}

func (ch *ContactHandler) HandleSubmit(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	var req worker.ContactPayload
	r.Body = http.MaxBytesReader(w, r.Body, 5120)

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid payload format or size limit exceeded"})
		return
	}

    if err := ch.validate.Struct(req); err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(map[string]string{"error": "Validation failed: please check your input data"})
        return
    }

    task, err := worker.NewContactTask(req.Name, req.Email, req.Message)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error creating task"})
        return
    }

    if _, err := ch.queueClient.Enqueue(task, asynq.MaxRetry(5)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error queuing task"})
        return
    }

    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}