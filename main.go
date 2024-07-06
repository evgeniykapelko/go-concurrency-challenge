package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage(p string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- p
}

func main() {
	var wg sync.WaitGroup

	phrase := []string{
		"Hello, universe",
		"Hello, cosmos!",
		"Hello World!",
	}

	wg.Add(len(phrase))
	ch := make(chan string, len(phrase))

	for _, p := range phrase {
		//updateMessage(p)
		go printMessage(p, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
