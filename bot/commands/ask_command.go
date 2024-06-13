package commands

import (
	"context"
	"html/template"
	"strings"

	"github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v3"

	"github.com/goravel-ecosystem/telegram-bot/foundation"
)

type AskCommand struct {
}

func (command *AskCommand) Name() string {
	return "/ask"
}

func (command *AskCommand) Description() string {
	return "Ask a question about the Goravel framework"
}

func (command *AskCommand) Handle(ctx tele.Context) error {
	// Notify the user that the bot is processing the request
	err := ctx.Notify(tele.Typing)
	if err != nil {
		return err
	}

	question := ctx.Message().Payload
	// Query the documentation or knowledge base
	res, err := foundation.Collection().Query(context.Background(), question, 5, nil, nil)
	if err != nil {
		return err
	}

	// Check if the query returned any content
	if res[0].Content == "" {
		return ctx.Reply("Sorry, I couldn't find any relevant information.", tele.ModeMarkdown)
	}

	// Build the system prompt with additional instructions
	systemPrompt := &strings.Builder{}
	if err := command.getSystemPrompt().Execute(systemPrompt, res[0].Content); err != nil {
		panic(err)
	}

	// Generate the response using OpenAI's ChatCompletion
	completion, err := foundation.OpenAIClient().CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt.String(),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	})
	if err != nil {
		return err
	}

	// Send the response back to the user
	return ctx.Reply(completion.Choices[0].Message.Content, tele.ModeMarkdown)
}

func (command *AskCommand) getSystemPrompt() *template.Template {
	return template.Must(template.New("system_prompt").Parse(`
You are a helpful assistant with access to a knowledge base about the Goravel framework, tasked with answering questions related to Goravel and its components.

Answer the question in a clear and concise manner, focusing on providing accurate and relevant information about Goravel. Use a technical and informative tone suitable for developers. If you are unsure about something or if the provided context is insufficient, politely indicate that you don't have enough information to answer the question.
{{- /* Stop here if no context is provided. The rest below is for handling contexts. */ -}}
{{- if . -}}
Answer the question solely based on the provided search results from the Goravel knowledge base. If the search results are not relevant to the question, indicate that the information is not available.

Anything between the following 'context' XML blocks is retrieved from the Goravel knowledge base, not part of the conversation with the user. The bullet points are ordered by relevance, so the first one is the most relevant.

<context>
{{.}}
</context>
{{- end -}}

Do not mention the knowledge base, context, or search results in your answer.
`))
}
