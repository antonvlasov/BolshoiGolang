package main

import (
	"HW-pt1/internal/pkg/storage"
	"fmt"
)

func main() {
	store, err := storage.InitStorage()
	if err != nil {
		fmt.Println(err)
		return
	}
	store.Set("string", "string")
	store.Set("int", "12345")
	store.Set("float64", "1234.123456789")
	store.Set("string with num", "234f")

	fmt.Println(store.Get("nil"))              // <nil>
	fmt.Println(*store.Get("string"))          // string
	fmt.Println(*store.Get("int"))             // 12345
	fmt.Println(*store.Get("float64"))         // 1234.123456789
	fmt.Println(*store.Get("string with num")) // 234f

	fmt.Println(store.GetKind("string"))          // S
	fmt.Println(store.GetKind("int"))             // D
	fmt.Println(store.GetKind("float64"))         // D
	fmt.Println(store.GetKind("nil"))             // N
	fmt.Println(store.GetKind("string with num")) // S

}
