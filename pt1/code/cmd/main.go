package main

import (
	"fmt"
	"livecode/internal/pkg/storage"
)

func main() {
	storage, err := storage.NewStorage()
	if err != nil {
		fmt.Println(err)
	}
	storage.Set("key", "value")
	fmt.Println(storage.Get("key"))
	x := any(5)
	fmt.Println(x)
	x = any("string")
	fmt.Println(x)
	xt := storage.GetKind("key")
	fmt.Println(xt)
	storage.Set("key2", "123")
	xt2 := storage.GetKind("key2")
	fmt.Println(xt2)
	test_nil := storage.Get("key4")
	if test_nil == nil {
		fmt.Println("nil")
	}
}
