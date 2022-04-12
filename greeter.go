package main

import "fmt"

func main() {
	fmt.Println(greet("Adam", "Welcome to Paradise on earth!"))
}

func greet(name, message string) string {
	return fmt.Sprintf("Hi %s! %s\n", name, message)
}
