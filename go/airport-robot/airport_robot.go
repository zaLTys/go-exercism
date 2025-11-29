package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}

func (italian Italian) LanguageName() string     { return "I can speak Italian: " }
func (italian Italian) Greet(name string) string { return fmt.Sprintf("Ciao %v!", name) }

type Portuguese struct{}

func (port Portuguese) LanguageName() string     { return "I can speak Portuguese: " }
func (port Portuguese) Greet(name string) string { return fmt.Sprintf("Olá %v!", name) }

func SayHello(name string, g Greeter) string { return g.LanguageName() + g.Greet(name) }
