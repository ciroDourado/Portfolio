package worker

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const TypeContactEmail = "email:contact"

type ContactPayload struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func NewContactTask(name, email, message string) (*asynq.Task, error) {
	payload, err := json.Marshal(ContactPayload{Name: name, Email: email, Message: message})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeContactEmail, payload), nil
}