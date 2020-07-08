package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}
func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sp spanishBot) {
// 	fmt.Println(sp.getGreeting())
// }
