package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	current := time.Now()
	exact, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	fmt.Printf("current time: %v \n", current)

	fmt.Printf("exact time: %v \n", exact)
	if err != nil {
		log.Fatal(err)
	}
}
