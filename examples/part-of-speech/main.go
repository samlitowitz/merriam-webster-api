package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/samlitowitz/merriam-webster-api/pkg/merriam-webster-api/collegiate/dictionary"
	"log"
	"net/http"
	"time"
)

func main() {
	var apiKey, userAgent, word string

	flag.StringVar(&apiKey, "api-key", "", "API Key with access to Merriam-Websters Collegiate Dictionary")
	flag.StringVar(&userAgent, "user-agent", "part-of-speech", "User-Agent for HTTP requests")
	flag.StringVar(&word, "word", "", "word to get the part of speech for")
	flag.Parse()

	if len(word) == 0 {
		log.Fatalf("Please enter a word")
	}

	httpClient := &http.Client{Timeout: time.Second * 10}
	client, err := dictionary.NewClient(apiKey, userAgent, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	var ctx context.Context
	responses, err := client.Lookup(ctx, word)
	if err != nil {
		log.Fatal(err)
	}

	for _, resp := range responses {
		fmt.Printf("%s, %s\n", resp.HeadwordInformation.Headword, resp.FunctionLabel)
	}
}
