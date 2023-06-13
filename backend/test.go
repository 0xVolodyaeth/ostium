package main

import (
	"time"
)

var done = make(chan struct{})

func main() {
	var ch chan int

	// go func() {
	// 	<-ch
	// }()

	ch <- 1

	// answers := sender()

	// go func() {
	// 	for {
	// 		select {
	// 		case answer, ok := <-answers:
	// 			fmt.Println(answer, ok)

	// 		}
	// 	}
	// }()

	// sigint := make(chan os.Signal, 1)
	// signal.Notify(sigint, os.Interrupt)
	// signal.Notify(sigint, syscall.SIGTERM)

	// <-sigint
	// done <- struct{}{}
}

func sender() chan int {

	answers := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case answers <- i:
			case <-done:
				close(answers)
				return

			}
			time.Sleep(time.Second * 3)
		}
	}()

	return answers
}
