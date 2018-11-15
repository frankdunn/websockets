package main

import (
	"strconv"
	"time"
)

func myfunc(broadcast chan []byte) {
	var count int
	for {
		broadcast <- []byte("hello " + strconv.Itoa(count))
		time.Sleep(time.Millisecond * 10)
		count++
	}
}
