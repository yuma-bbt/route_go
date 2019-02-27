package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
}

func (u user) show() {
	fmt.Printf("name: %s, score: %d\n", u.name, u.score)
}

func (u *user) hit() {
	u.score++
}

func main() {
	u := user{name: "fujimoto", score: 500}
	u.hit()
	u.show()
}
