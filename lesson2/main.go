package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) Person {
	return Person{Name: name, Age: 0}
}

func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("Happy birthday, %s! You are now %d years old.\n", p.Name, p.Age)
}

type Book struct {
	Title  string
	Author string
}

type LibraryCard struct {
	CardNumber string
	Owner      Person
	Book       Book
}

func NewBook(title, author string) *Book {
	return &Book{Title: title, Author: author}
}

func NewLibraryCard(cardNumber string, owner Person, book Book) *LibraryCard {
	return &LibraryCard{CardNumber: cardNumber, Owner: owner, Book: book}
}

func (b *Book) Info() string {
	return fmt.Sprintf("'%s' by %s", b.Title, b.Author)
}

type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Address Address
}

func FullAddress(u User) string {
	return fmt.Sprintf("%s, %s", u.Address.City, u.Address.Street)
}

func main() {
	fmt.Println("--- Person")
	p1 := NewPerson("Vladimir")
	fmt.Println("До дня рождения:", p1)
	p1.Birthday()
	fmt.Println("После дня рождения:", p1)

	fmt.Println("\n--- Book")
	b1 := Book{"Go Basics", "John Doe"}
	fmt.Println("Литерал ", b1)

	b2 := NewBook("The Go Programming Language", "Alan A. A. Donovan")
	fmt.Println("Конструктор:", *b2)

	b3 := new(Book)
	fmt.Println("new(Book)", *b3)

	fmt.Println("\n--- LibraryCard")
	book := NewBook("1984", "George Orwell")
	owner := NewPerson("Vova")
	card := NewLibraryCard("1", owner, *book)
	fmt.Printf("Library Card: %s owns %s\n", card.Owner.Name, card.Book.Info())

	fmt.Println("\n--- User Address")
	user := User{
		Name: "Alice",
		Address: Address{
			City:   "Almaty",
			Street: "Abai Ave 10",
		},
	}
	fmt.Println("User: ", user)
	fmt.Println("FullAddress:", FullAddress(user))

}
