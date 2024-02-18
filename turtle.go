package main

import "fmt"

func turtle() error {
	fmt.Println(talk())
	return nil
}

func talk() string {
	return "I'm walking slowly."
}
