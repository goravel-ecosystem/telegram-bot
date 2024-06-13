package config

import "github.com/goravel-ecosystem/telegram-bot/foundation"

func init() {
	config := foundation.Config()

	config.Set("openai", map[string]any{
		"api_key": config.GetString("OPENAI_API_KEY"),
	})
}
