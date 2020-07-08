package main

type user struct {
	name string
}

type bots interface {
	getGreetings(string, int) (string, error)
	getBotVersions() string
	respondToUser(user) string
}
