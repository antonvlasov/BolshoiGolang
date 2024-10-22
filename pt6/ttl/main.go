package main

import (
	"fmt"
	"log"
	"time"
)

var storage map[string]value

type value struct {
	v         any
	expiresAt int64
}

func main() {
	storage := make(map[string]value)

	ttl := 10 * time.Second

	storage["1"] = value{
		v:         "a12lksdfjkl",
		expiresAt: time.Now().Add(ttl).UnixMilli(),
	}

	v, ok := storage["1"]
	if !ok {
		log.Fatal("")
	}

	if time.Now().UnixMilli() >= v.expiresAt {
		fmt.Println("expired")
		return
	}

	closeChan := make(chan struct{})
	go iWantToSleepFor(closeChan, time.Minute*10)

	close(closeChan)

	fmt.Println(v.v)
}

func iWantToSleepFor(closeChan chan struct{}, n time.Duration) {
	for {
		select {
		case <-closeChan:
			return
		case <-time.After(n):
			Clean()
		}
	}
}

func Clean() {
	// Lock
	// for
	// if expired - delete
	// Unlock
}
