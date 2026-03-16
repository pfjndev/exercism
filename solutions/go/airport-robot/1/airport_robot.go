package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
import "fmt"

type Italian struct {
	greet			string
	languageName	string
}

func (it Italian) Greet(name string) string {
	it.greet = "Ciao"
	return fmt.Sprintf("%s %s!", it.greet, name)
}


func (it Italian) LanguageName() string {
	it.languageName = "Italian"
	return fmt.Sprintf("%s", it.languageName)
}

type Portuguese struct {
	greet			string
	languageName	string
}

func (pt Portuguese) Greet(name string) string {
	pt.greet = "Olá"
	return fmt.Sprintf("%s %s!", pt.greet, name)
}


func (pt Portuguese) LanguageName() string {
	pt.languageName = "Portuguese"
	return fmt.Sprintf("%s", pt.languageName)
}

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}