package main

import (
	"fmt"
)

func Hello(p ...string) (int, error) {
	fmt.Printf("Hello: %v\n", p)
	return len(p), nil
}

func Say(p string) int {
	fmt.Printf("You say: %s\n", p)
	return len(p)
}
