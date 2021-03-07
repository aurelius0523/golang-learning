package main

import "fmt"

func main() {
	hello()
}

func hello() (int, error) {
	return fmt.Println("Hello", "world")
}

