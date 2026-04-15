package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/cirodourado/portfolio/applications/api/internal/models"
	"github.com/cirodourado/portfolio/applications/api/internal/notifications"
	"github.com/cirodourado/portfolio/applications/api/internal/worker"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type WorkerHandler struct {
	db          *gorm.DB
	queueClient *asynq.Client
	logger      *slog.Logger
}

func NewWorkerHandler(db *gorm.DB, qc *asynq.Client, logger *slog.Logger) *WorkerHandler {
	return &WorkerHandler{
		db:          db,
		queueClient: qc,
		logger:      logger,
	}
}

func (wh *WorkerHandler) HandleCoordinator(ctx context.Context, t *asynq.Task) error {
	var payload worker.ContactPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		wh.logger.Error("Failed to parse payload", "error", err)
		return fmt.Errorf("json parsing failed: %v", err)
	}

	msg := models.ContactMessage{
		Name:    payload.Name,
		Email:   payload.Email,
		Message: payload.Message,
		Status:  "processing",
	}

	if err := wh.db.Create(&msg).Error; err != nil {
		wh.logger.Error("Database insert failed", "error", err)
		return err
	}

	taskDiscord, _ := worker.NewDiscordTask(payload)
	taskTelegram, _ := worker.NewTelegramTask(payload)

	_, errD := wh.queueClient.EnqueueContext(ctx, taskDiscord)
	_, errT := wh.queueClient.EnqueueContext(ctx, taskTelegram)

	if errD != nil || errT != nil {
		wh.db.Model(&msg).Update("status", "failed")
		wh.logger.Error("Fan-out failed", "discordErr", errD, "telegramErr", errT)
		return fmt.Errorf("failed to enqueue child tasks")
	}

	wh.db.Model(&msg).Update("status", "enqueued")
	wh.logger.Info("Coordinator finished successfully", "email", payload.Email)

	return nil
}

func (wh *WorkerHandler) HandleDiscord(ctx context.Context, t *asynq.Task) error {
	var payload worker.ContactPayload
	json.Unmarshal(t.Payload(), &payload)

	if err := notifications.SendDiscord(payload); err != nil {
		wh.logger.Error("Discord delivery failed", "error", err)
		return err
	}

	wh.logger.Info("Discord delivery successful", "email", payload.Email)
	return nil
}

func (wh *WorkerHandler) HandleTelegram(ctx context.Context, t *asynq.Task) error {
	var payload worker.ContactPayload
	json.Unmarshal(t.Payload(), &payload)

	if err := notifications.SendTelegram(payload); err != nil {
		wh.logger.Error("Telegram delivery failed", "error", err)
		return err
	}

	wh.logger.Info("Telegram delivery successful", "email", payload.Email)
	return nil
}