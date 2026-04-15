package main

import (
	"log/slog"
	"os"

	"github.com/cirodourado/portfolio/applications/api/internal/config"
	"github.com/cirodourado/portfolio/applications/api/internal/database"
	"github.com/cirodourado/portfolio/applications/api/internal/handlers"
	"github.com/cirodourado/portfolio/applications/api/internal/worker"
	"github.com/hibiken/asynq"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadWorkerConfig()
	if err != nil {
		logger.Error("Erro ao carregar configurações", "error", err)
	}

	db, err := database.NewConnection(cfg.GetDatabaseURL())
	if err != nil {
		logger.Error("Erro ao conectar no banco", "error", err)
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.GetRedisURL(),
	}

	queueClient := asynq.NewClient(redisOpt)
	defer queueClient.Close()

	queueServer := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 1,
			},
		},
	)

	workerHandler := handlers.NewWorkerHandler(db, queueClient, logger)

	mux := asynq.NewServeMux()
	mux.HandleFunc(worker.TypeContactProcess, workerHandler.HandleCoordinator)
	mux.HandleFunc(worker.TypeContactDiscord, workerHandler.HandleDiscord)
	mux.HandleFunc(worker.TypeContactTelegram, workerHandler.HandleTelegram)

	logger.Info("Worker started")
	if err := queueServer.Run(mux); err != nil {
		logger.Error("Could not run server", "error", err)
		os.Exit(1)
	}
}