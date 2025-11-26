package main

import "fmt"

type Saver interface {
	Save() error
}

type Config struct {
	Host  string
	Port  string
	User  string
	Pass  string
	Saved bool
}

func (c Config) Save() error {
	fmt.Println(fmt.Sprintf("%s:%s@%s:%s", c.User, c.Pass, c.Host, c.Port))
	c.Saved = true
	return nil
}

func main() {
	var dbConfig Saver = &Config{
		Host: "localhost",
		Port: "5432",
		User: "asd",
		Pass: "asd",
	}
	err := dbConfig.Save()
	if err != nil {
		fmt.Println(err)
	}
}
