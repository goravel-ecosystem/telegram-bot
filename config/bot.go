package config

import "github.com/goravel-ecosystem/telegram-bot/foundation"

func init() {
	config := foundation.Config()

	config.Set("bot", map[string]any{
		"token": config.GetString("TELEGRAM_BOT_TOKEN"),

		"poller": map[string]any{
			"timeout": 10,
		},
	})
}
