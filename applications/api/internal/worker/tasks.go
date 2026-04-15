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
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
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