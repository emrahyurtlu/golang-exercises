package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{FirstName: "John", LastName: "Simith", Email: "john@simith.com", Age: 30})
	users = append(users, User{FirstName: "Mary", LastName: "Abel", Email: "mary@simith.com", Age: 20})
	users = append(users, User{FirstName: "Elise", LastName: "Calinda", Email: "elis@simith.com", Age: 40})
	users = append(users, User{FirstName: "Mathilda", LastName: "Stew", Email: "math@simith.com", Age: 25})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}
