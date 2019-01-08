package main

import (
	"fmt"

	lru "github.com/harleylau/go-lru"
)

func main() {
	l, _ := lru.NewLRU(128)
	for i := 0; i < 256; i++ {
		l.Put(i, i)
	}
	res, err := l.Get(45)
	if err == nil {
		println(fmt.Sprintf("get value: %v", res))
	} else {
		println(fmt.Sprintf("get err: %v", err))
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
}
