package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello %s", u.Name)
}

func PrintGreeting(g Greeter) {
	u, ok := g.(User)
	if !ok {
		fmt.Println("Greeter is a nil *User")
		return
	}
	fmt.Println(u.Greet())
}

func PrintType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Type is int with value", v)
	case string:
		fmt.Println("Type is string with value", v)
	case User:
		fmt.Println("Type is User with name", v.Name)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var someValue interface{} = 123

	v, ok := someValue.(int)
	if ok {
		fmt.Println("Value is an int:", v)
	} else {
		println("Value is not an int")
	}

	s, ok := someValue.(string)
	if ok {
		fmt.Println("Value is a string:", s)
	} else {
		println("Value is not a string")
	}

	var u User = User{Name: "Alice"}
	PrintGreeting(u)

	PrintType(u)

	// var u2 User = nil
	// PrintGreeting(u2)

}
