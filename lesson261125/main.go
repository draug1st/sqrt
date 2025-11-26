package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

type Robot struct {
	Model string
}

func (r Robot) Greet() string {
	return "Beep boop, I am model " + r.Model
}

func SayHello(g Greeter) string {
	return g.Greet()
}

func main() {

	var SomeGreeter Greeter
	fmt.Println(SomeGreeter)
	println(SomeGreeter)

	andrew := User{Name: "Andrew"}
	robot := Robot{Model: "RX-78"}
	SomeGreeter = User{Name: "Alice"}

	println(SayHello(andrew))          // Output: Hello, Andrew
	println(SayHello(robot))           // Output: Beep boop, I am model RX-78
	fmt.Println(SayHello(SomeGreeter)) // Output: Hello, Alice
}
