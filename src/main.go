package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	channel := make(chan string)
	doneChannel := make(chan bool)

	go letterGenerator(channel)
	go printer(channel, doneChannel)

	println("Stuff is happening")
	<-doneChannel
}

func letterGenerator(ch chan string) {
	for letter := byte('a'); letter <= byte('z'); letter++ {
		ch <- string(letter)
	}

	close(ch)
}

func printer(ch chan string, done chan bool) {
	for letter := range ch {
		println(letter)
	}

	done <- true
}
