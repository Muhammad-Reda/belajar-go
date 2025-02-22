package channel

import (
	"fmt"
	"time"
)

func portal1(chanel1 chan string) {
	// time.Sleep(3 * time.Second)
	chanel1 <- "From portal 1"
}

func portal12(channel2 chan string) {
	// time.Sleep(3 * time.Second)
	channel2 <- "From portal 2"
}

func BasicChan() {
	ch := make(chan string)
	ch2 := make(chan string)

	// go func() {
	// 	fmt.Println("Mengirim pesan")
	// 	ch <- "Hello dari goroutine"
	// 	fmt.Println("Pesan terkirim")
	// }()

	// msg := <-ch
	// fmt.Println("Pesan diterima ", msg)

	go portal1(ch)
	go portal12(ch2)

	time.Sleep(2 * time.Second)
	select {
	case op1 := <-ch:
		fmt.Println(op1)
	case op2 := <-ch2:
		fmt.Println(op2)
	default:
		fmt.Println("No case ready yet")
	}
}
