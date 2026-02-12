package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(visitorName string) string
}

func SayHello(name string, greeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {}

func (italian Italian) LanguageName() string {
    return "Italian"
}

func (italian Italian) Greet(visitorName string) string {
    return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct {}

func (portugese Portuguese) LanguageName() string {
    return "Portuguese"
}

func (portugese Portuguese) Greet(visitorName string) string {
    return fmt.Sprintf("Olá %s!", visitorName)
}