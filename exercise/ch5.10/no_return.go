package main

import "fmt"

func main() {
	defer func() {
		switch p := recover(); p.(type) {
		case int64:
			fmt.Println(p.(int64))
		default:
			panic(p)
		}
	}()
	noReturn(10)
}

func noReturn(num int64) {
	panic(num)
}
