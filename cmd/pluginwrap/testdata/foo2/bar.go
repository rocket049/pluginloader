package main

import (
	"fmt"
	"time"
)

func Bar() {
	fmt.Println(time.Now())
}

type little struct{}

func (s *little) Hello() {
	fmt.Println("hello")
}

func NewLittle() (*little, error) {
	return new(little), nil
}

func main() {}
