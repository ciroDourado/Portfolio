package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cirodourado/portfolio/applications/api/internal/worker"
)

type DiscordEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
}

type DiscordMessage struct {
	Embeds []DiscordEmbed `json:"embeds"`
}

func SendDiscord(payload worker.ContactPayload) error {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL is not set")
	}

	msg := DiscordMessage{
		Embeds: []DiscordEmbed{
			{
				Title:       fmt.Sprintf("Novo Contato: %s", payload.Name),
				Description: fmt.Sprintf("**Email:** %s\n\n**Mensagem:**\n%s", payload.Email, payload.Message),
				Color:       3447003,
			},
		},
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("discord API returned status: %d", resp.StatusCode)
	}

	return nil
}