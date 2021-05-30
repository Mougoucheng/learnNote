package Command

import (
	"log"
)

type Command struct {
	person *Person
	method func()
}

func NewCommand(p *Person, method func()) Command {
	return Command{
		person: p,
		method: method,
	}
}

func (th *Command) Execute() {
	th.method()
}

type Person struct {
	name string
	cmd  Command
}

func NewPerson(name string, cmd Command) Person {
	return Person{
		name: name,
		cmd:  cmd,
	}
}

func (th *Person) Buy() {
	log.Printf("%s is Buying", th.name)
	th.cmd.Execute()
}

func (th *Person) Cook() {
	log.Printf("%s is Cooking", th.name)
	th.cmd.Execute()
}

func (th *Person) Wash() {
	log.Printf("%s is Washing", th.name)
	th.cmd.Execute()
}

func (th *Person) Listen() {
	log.Printf("%s is Listening", th.name)
	//th.cmd.Execute()
}

func (th *Person) Talk() {
	log.Printf("%s is Talking", th.name)
	th.cmd.Execute()
}
