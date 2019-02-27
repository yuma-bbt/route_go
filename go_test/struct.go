package main

import "fmt"

type user struct {
	name  string
	score int
}

func main() {
	u1 := new(user)

	u1.name = "fujimoto"
	u1.score = 200

	fmt.Println(u1)

	u2 := user{name: "arita", score: 300}
	fmt.Println(u2)
}
