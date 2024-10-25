package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func consumeResponse(r io.ReadCloser) {
	_, err := io.Copy(io.Discard, r)
	if err != nil {
		log.Print(err)
	}

	if err := r.Close(); err != nil {
		log.Print(err)
	}
}

func main() {
	serverAddress, ok := os.LookupEnv("BASIC_CLIENT_SERVER_ADDRESS")
	if !ok {
		fmt.Println("server address not provided")
		os.Exit(1)
	}

	//
	// Init client.
	//

	cli := http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//
	// GET request.
	//

	healthReq, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s/health", serverAddress), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Do(healthReq)
	if err != nil {
		log.Fatal(err)
	}

	defer consumeResponse(resp.Body)

	fmt.Println(resp.StatusCode)
}
