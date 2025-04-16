package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(time)
}
