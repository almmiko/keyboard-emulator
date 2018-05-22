package main

import (
	"os"
	"fmt"
	"time"
	"flag"
	"github.com/go-vgo/robotgo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scanner-emulator -h")
		os.Exit(1)
	}

	barcode := flag.String("barcode", "", "Barcode (Required)")
	timeout := flag.Int("timeout", 0, "Timeout in milliseconds (Optional)")

	flag.Parse()

	if *timeout > 0 {
		time.Sleep(time.Millisecond * time.Duration(*timeout))
	}

	robotgo.TypeString(*barcode)

	time.Sleep(time.Millisecond * 1000)
	os.Exit(0)
}
