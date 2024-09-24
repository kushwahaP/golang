package main

import (
	"fmt"
	"strconv"
	"time"
)

func test_channel() {

	message := make(chan string)

	go func() {
		message <- "data send to channel"
	}()

	msg := <-message
	fmt.Println(msg)
}

func test_channel_concurrency() {
	ch := make(chan string)

	go countfun("channel data", ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func countfun(data string, ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- data + " " + strconv.Itoa(i)
	}
	close(ch)
}

func ping(pings chan<- string, msg string) {
	fmt.Println("inside the ping function")
	pings <- msg

}

func pong(pings <-chan string, pongs chan<- string) {
	fmt.Println("inside the pong functions")
	msg := <-pings
	pongs <- msg

}

func test_channel_directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "send ping data")
	fmt.Println("inside the main function")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func test_select_with_channel() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		channel1 <- "channel one"
	}()
	go func() {
		time.Sleep(time.Second)
		channel2 <- "channel two"
	}()

	for i := 0; i < 2; i++ {

		select {
		case msg1 := <-channel1:
			fmt.Println("changes one value ", msg1)
		case msg2 := <-channel2:
			fmt.Println("channel tow value :", msg2)

		}

	}
}

func timeout_example_channel() {
	channel1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		channel1 <- "testing result with the first channel"
	}()

	select {
	case resp := <-channel1:
		fmt.Println(resp)

	case <-time.After(time.Second * 2):
		fmt.Println("response time out for the channel one ")

	}

	channel2 := make(chan string, 2)
	go func() {
		time.Sleep(time.Second * 3)
		channel2 <- "testing result with the  second channel"
	}()

	select {
	case resp := <-channel2:
		fmt.Println(resp)
	case <-time.After(time.Second * 4):
		fmt.Println("response time out for channel two")
	}
}
