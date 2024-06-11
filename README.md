# telegram-bot

Telegram bot for answering questions about the Goravel framework.

## Installation

```bash
git clone github.com/goravel-ecosystem/telegram-bot
```

## Development

```bash
cd telegram-bot

# copy .env.example to .env
cp .env.example .env

# run the bot
go run main.go
```

## Directory Structure

The project is organized into several directories, each serving a specific purpose. Below is an overview of the directory structure:

### /routes

Contains the `bot.go` file, which is responsible for defining handlers that are not command-specific.

### /config

Contains the `bot.go` file, which is used for configuring the bot.

### /bot/commands

Contains the implementation of various commands. Each command should implement the interface defined in `contracts/command/command.go`.

### /bot/handlers

Contains custom handlers that can be used in `routes/bot.go`.

### /bot/kernel.go

Contains the `Kernel` struct, which is responsible for registering commands.
