# Go client for Merriam-Websters Dictionary API
[![Go Report Card](https://goreportcard.com/badge/github.com/samlitowitz/merriam-webster-api)](https://goreportcard.com/report/github.com/samlitowitz/merriam-webster-api)
[![GoDoc](https://godoc.org/github.com/samlitowitz/merriam-webster-api/pkg/merriam-webster-api/collegiate/dictionary?status.svg)](https://godoc.org/github.com/samlitowitz/merriam-webster-api/pkg/merriam-webster-api/collegiate/dictionary)

# Install
```bash
go get -u github.com/samlitowitz/merriam-webster-api
```

# Getting an API key
Register with Merriam-Websters developer center at [https://dictionaryapi.com](https://dictionaryapi.com).
Currently only the Collegiate Dictionary is supported.

# Using the Collegiate Dictionary package
```go
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
	apiKey, userAgent, word := "your-api-key", "a-merry-user-agent", "test" 
	
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
```
