package main

import (
	"fmt"
	"time"
)

func Help(s string) (a string, e error) {
	fmt.Println(s, time.Now())
	return s, nil
}

type Man struct{}

func (s *Man) Hello() error {
	fmt.Println("hello")
	return nil
}

func NewMan() *Man {
	return new(Man)
}
