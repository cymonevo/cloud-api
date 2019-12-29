package main

import (
	"net/http"

	"github.com/cymonevo/cloud-api/internal/log"
	"github.com/cymonevo/cloud-api/provider"
	"github.com/cymonevo/cloud-api/sample"
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
	sample.KloudlessSample()
}

func Serve() {
	router := provider.GetArticleHandler().Register()
	log.FatalDetail("MAIN", "abort server", http.ListenAndServe(":8000", router))
}
