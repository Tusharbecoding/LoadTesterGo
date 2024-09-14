package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomSleep() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000) // Random delay between requests
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println("Slept for", n, "milliseconds")
}
