package main

import (
    "context"

	"github.com/hibiken/asynq"
	"github.com/cirodourado/portfolio/applications/api/internal/worker"
)

func main() {
    //
}

func HandleContactEmail(ctx context.Context, t *asynq.Task) error {
    return nil
}

func sendDiscord(p worker.ContactPayload) error {
    return nil
}

func sendTelegram(p worker.ContactPayload) error {
    return nil
}