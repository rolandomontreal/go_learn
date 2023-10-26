package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}
// 	printGreeting(eb)
// 	printGreeting(sb)
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello there INGLES"
}

func (sb spanishBot) getGreeting() string {
	return "HOLA espanol"
}