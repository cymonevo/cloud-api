package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cymonevo/cloud-api/internal/log"
	"github.com/cymonevo/cloud-api/provider"
)

func main() {
	//consumers := provider.GetArticleConsumers()
	//for _, c := range consumers {
	//	c.Consume()
	//}
	Client()
	//Serve()
}

func Client() {
	client := provider.GetKloudlessClient()
	resp, err := client.GetAccounts(context.Background())
	fmt.Println(resp, err)
}

func Serve() {
	router := provider.GetArticleHandler().Register()
	log.FatalDetail("MAIN", "abort server", http.ListenAndServe(":8000", router))
}
