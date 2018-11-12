package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	min := 100000000
	max := 999999999

	log.Println(getRand(min, max))
}

func getRand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
