package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

	"github.com/hibiken/asynq"
	"github.com/cirodourado/portfolio/applications/api/internal/config"
	"github.com/cirodourado/portfolio/applications/api/internal/worker"
)

func main() {
    cfg, err := config.LoadWorkerConfig()
	if err != nil {
		log.Fatalf("Error loading configs: %v", err)
	}

	queueServer := asynq.NewServer(
		asynq.RedisClientOpt{
		    Addr: cfg.GetRedisURL(),
        },
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 1,
			},
		},
	)

    mux := asynq.NewServeMux()
    mux.HandleFunc(worker.TypeContactEmail, HandleContactEmail)

    if err := queueServer.Run(mux); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}

func HandleContactEmail(ctx context.Context, t *asynq.Task) error {
    var p worker.ContactPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json parsing failed: %v", err)
	}

    if err := sendDiscord(p); err != nil {
		return err
	}

	if err := sendTelegram(p); err != nil {
		return err
	}

    return nil
}

func sendDiscord(p worker.ContactPayload) error {
    return nil
}

func sendTelegram(p worker.ContactPayload) error {
    return nil
}