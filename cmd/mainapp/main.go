package main

import (
	"net/http"

	"github.com/cymonevo/cloud-api/internal/log"
	"github.com/cymonevo/cloud-api/provider"
)

func main() {
	consumers := provider.GetArticleConsumers()
	for _, c := range consumers {
		c.Consume()
	}

	handlers := provider.GetArticleHandler()
	log.Fatalf("Aborting...", http.ListenAndServe(":8000", handlers.Register()))
}
