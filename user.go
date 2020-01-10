package main

import (
	"fmt"
)

type User struct {
	Age int
	Context
}

func (u User) GetUser() {
	fmt.Print("success")
	fmt.Print(u.Age)
}

func init() {
	RegisterType(User{})
}
