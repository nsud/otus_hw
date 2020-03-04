package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const form = "2006-01-02 15:04:05 +0000 UTC"
	current := time.Now()
	exact, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	fmt.Printf("current time: %v\n", current.Format(form))

	fmt.Printf("exact time: %v", exact.Format(form))
	if err != nil {
		log.Fatal(err)
	}
}
