package main

import (
	"context"
	"fmt"

	"github.com/philippgille/chromem-go"

	"github.com/goravel-ecosystem/telegram-bot/support"
)

var OpenaiApiKey = ""

func createEmbedding() {
	repoPath := "../docs"
	docs, err := support.ReadDocs(repoPath)
	if err != nil {
		fmt.Printf("Error reading docs: %v\n", err)
		return
	}

	ctx := context.Background()

	db, err := chromem.NewPersistentDB("", true)
	if err != nil {
		panic(err)
	}

	c, err := db.CreateCollection("goravel_bot", nil, chromem.NewEmbeddingFuncOpenAI(OpenaiApiKey, chromem.EmbeddingModelOpenAI3Small))
	if err != nil {
		panic(err)
	}

	id := 0
	for path, content := range docs {
		err = c.AddDocument(ctx, chromem.Document{
			ID:      fmt.Sprintf("%d", id),
			Content: content,
		})
		if err != nil {
			fmt.Printf("Error adding document %s: %v\n", path, err)
		}
		id++
	}
}
