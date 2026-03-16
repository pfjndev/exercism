package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Italian struct {
	greet			string
	languageName	string
}

func (it Italian) Greet(name string) string {
	it.greet = "Ciao"
	return it.greet + " " + name + "!"
}


func (it Italian) LanguageName() string {
	it.languageName = "Italian"
	return it.languageName
}

type Portuguese struct {
	greet			string
	languageName	string
}

func (pt Portuguese) Greet(name string) string {
	pt.greet = "Olá"
	return pt.greet + " " + name + "!"
}


func (pt Portuguese) LanguageName() string {
	pt.languageName = "Portuguese"
	return pt.languageName
}

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}