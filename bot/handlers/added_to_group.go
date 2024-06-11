package handlers

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

type AddedToGroupHandler struct{}

func NewAddedToGroupHandler() *AddedToGroupHandler {
	return &AddedToGroupHandler{}
}

func (handler *AddedToGroupHandler) Handle(ctx tele.Context) error {
	// Notify the user that the bot is typing
	if err := ctx.Notify(tele.Typing); err != nil {
		return err
	}

	// Prepare the greeting message
	message := fmt.Sprintf("Hello, *%s*! 👋\n", ctx.Message().Sender.FirstName)
	message += "Thanks for adding me! Here’s what I can do:\n\n"
	message += "*Features:*\n"
	message += "🔍 Search docs\n"
	message += "📈 Upgrade guides\n"
	message += "💻 Code snippets\n"
	message += "❓ FAQs\n\n"
	message += "Type `/help` to see commands."

	// Send the message with Markdown formatting
	return ctx.Send(message, tele.ModeMarkdown)
}
