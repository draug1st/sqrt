package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type IRepository interface {
	FindByName(name string) (*User, *UserNotFoundError)
	FindAll() []*User
}

type Repository struct {
	Data []*User
}

type UserNotFoundError struct {
	cause error
	Code  int
	Name  string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with name '%s' not found: %v", e.Name, e.cause)
}

func (e *UserNotFoundError) Unwrap() error {
	return e.cause
}

func (r *Repository) FindByName(name string) (*User, *UserNotFoundError) {
	for _, user := range r.Data {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, &UserNotFoundError{
		cause: errors.New("no such user in repository"),
		Code:  404,
		Name:  name,
	}
}

func (r *Repository) FindAll() []*User {
	return r.Data
}

func main() {
	a := []string{"one", "two", "three"}

	for i, v := range a {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	u := []User{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}

	for i, user := range u {
		if user.Name == "Bob" {
			user.Age += 1
			u[i] = user
		}
		fmt.Printf("User %d: Name=%s, Age=%d\n", i, user.Name, user.Age)
	}

	userRepo := Repository{
		Data: make([]*User, 0),
	}
	userRepo.Data = append(userRepo.Data, &User{Name: "Alice", Age: 10})
	userRepo.Data = append(userRepo.Data, &User{Name: "Bob", Age: 20})

	user, uerr := userRepo.FindByName("Bob")
	if uerr != nil {
		fmt.Println("Error:", uerr)
	} else {
		fmt.Printf("Found user: Name=%s, Age=%d\n", user.Name, user.Age)
	}

	fmt.Println("-------------------\nAll users in repository:")
	for _, u := range userRepo.FindAll() {
		fmt.Printf("Name=%s, Age=%d\n", u.Name, u.Age)
	}

	s := "Hello, World!"
	for i, ch := range s {
		fmt.Printf("i: %d, ch: %c\n", i, ch)
	}

	var m string
	for _, ch := range []rune{'п', 'р', 'и', 'в', 'е', 'т'} {
		m += string(ch)
	}
	fmt.Println("---------------------\n", m)

	x := map[int]map[string]int{
		1: {"one": 1, "two": 2},
		2: {"three": 3, "four": 4},
	}

	x[3] = map[string]int{"five": 5, "six": 6}

	delete(x, 1)

	fmt.Println(x)

	p, ok := x[4]
	fmt.Println(ok, p)

	j := `{"name":"Alice","age":30}`
	//result := make(map[string]interface{})
	result := Request{}

	err := json.Unmarshal([]byte(j), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.Name, result.Age)
	}

}

type Request struct {
	Name interface{} `json:"name"`
	Age  interface{} `json:"age"`
}
