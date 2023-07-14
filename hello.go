package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
