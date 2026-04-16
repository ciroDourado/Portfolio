package worker

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeContactProcess  = "contact:process"
	TypeContactDiscord  = "contact:discord"
	TypeContactTelegram = "contact:telegram"
)

type ContactPayload struct {
    Name    string `json:"name" validate:"required,max=100"`
	Email   string `json:"email" validate:"required,email,max=100"`
	Message string `json:"message" validate:"required,max=1000"`
}

func NewContactTask(name, email, message string) (*asynq.Task, error) {
	payload, err := json.Marshal(ContactPayload{Name: name, Email: email, Message: message})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeContactProcess, payload), nil
}

func NewDiscordTask(p ContactPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeContactDiscord, payload), nil
}

func NewTelegramTask(p ContactPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeContactTelegram, payload), nil
}