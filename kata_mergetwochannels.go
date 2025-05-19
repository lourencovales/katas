//Task

//Write a function func Merge(a <-chan string, b <-chan string) <-chan string, which takes two read-only channels and returns a new channel. All messages from channel a and b must be forwarded to the new channel. Once a and b are both closed, also the returned channel must be closed.

//The order of the forwarded messages doesn't matter, but you should consume from both incoming channels concurrenly.

package main

import (
	"sync"
)

func main() {

}

func Merge(a <-chan string, b <-chan string) <-chan string {
	return concMerge(a, b)
}

func concMerge(channel ...<-chan string) <-chan string {
	merge := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range channel {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				merge <- msg
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(merge)
	}()
	return merge
}
