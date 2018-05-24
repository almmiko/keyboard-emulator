package main

import (
	"flag"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scanner-emulator -h")
		os.Exit(1)
	}

	barcode := flag.String("barcode", "", "Barcode (Required)")
	timeout := flag.Int("timeout", 0, "Timeout in milliseconds (Optional)")

	flag.Parse()

	done := make(chan bool)

	if *timeout > 0 {
		time.Sleep(time.Millisecond * time.Duration(*timeout))
	}

	go typeString(*barcode, done)

	<-done

	os.Exit(0)
}

func typeString(barcode string, done chan<- bool) {
	robotgo.TypeString(barcode)
	done <- true
}
