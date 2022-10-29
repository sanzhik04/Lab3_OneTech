package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan<-string, msg string) {
	// send message on ch1
	ch1 <- msg
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	// send it on ch2
	msg:= <- ch1
	msg = msg + " *relayed*"
	ch2 <- msg 
}

func main() {
	// create ch1 and ch2
	ch1:= make(chan string)
	ch2:= make(chan string)
	// spine goroutine genMsg and relayMsg
	go func(){
		genMsg(ch1, "Hi, how are you?")
	}()

	go func(){
		relayMsg(ch1,ch2)
	}()


	// recv message on ch2

	fmt.Println(<-ch2)
}
