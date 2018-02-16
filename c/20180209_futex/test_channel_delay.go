package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("starting...")
	ch := make(chan bool, 1)
	go func() {
		<-ch
		log.Printf("%v", time.Now().UnixNano())
	}()

	time.Sleep(2 * time.Second)
	log.Printf("%v", time.Now().UnixNano())
	ch <- true
	time.Sleep(2 * time.Second)
}

// macos:
// $ go run test_channel_delay.go
// 2018/02/11 19:19:19 starting...
// 2018/02/11 19:19:21 1518347961360782000
// 2018/02/11 19:19:21 1518347961360941000
// lubuntu@hp
// $ ./test_channel_delay 
// 2018/02/12 19:14:52 starting...
// 2018/02/12 19:14:54 1518434094798942301
// 2018/02/12 19:14:54 1518434094799042050
// chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay 
// 2018/02/12 19:15:14 starting...
// 2018/02/12 19:15:16 1518434116993779883
// 2018/02/12 19:15:16 1518434116993938206
// 结论：大约100us左右
