package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
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
	//
	// Init client.
	//

	cli := http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//
	// GET request.
	//

	healthReq, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:7500/health", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Do(healthReq)
	if err != nil {
		log.Fatal(err)
	}

	defer consumeResponse(resp.Body)

	fmt.Println(resp.StatusCode)

	// //
	// // POST request.
	// //

	// x2Request := types.X2Request{
	// 	Val: 2,
	// }

	// b, err := json.Marshal(x2Request)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err = cli.Post("http://127.0.0.1:7500/x2", "", bytes.NewReader(b))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer consumeResponse(resp.Body)

	// var x2Resp types.X2Response
	// if err = json.NewDecoder(resp.Body).Decode(&x2Resp); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(x2Resp.Val)

	// fmt.Println(resp.StatusCode)
}
