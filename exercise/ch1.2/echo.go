package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func slowEcho() {
	for index, arg := range os.Args {
		fmt.Println(strconv.Itoa(index) + " : " + arg)
	}
}

func fastEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func measureTime(f func(), msg string) {
	now := time.Now()
	f()
	fmt.Printf("[%s] cost time: %d\n", msg, time.Since(now).Nanoseconds())
}

func main() {
	measureTime(slowEcho, "slow")
	measureTime(fastEcho, "fast")
}
