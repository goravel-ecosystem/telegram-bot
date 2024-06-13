package config

import "github.com/goravel-ecosystem/telegram-bot/foundation"

func init() {
	config := foundation.Config()

	config.Set("database", map[string]any{
		"collection": config.GetString("DATABASE_COLLECTION"),
	})
}
