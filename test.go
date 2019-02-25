package main

type Person struct {
	firstName string
	age       int
}

func main() {
	tim := Person{"Tim", 25}
	person1 := &tim
	(*person1).age = 25
	person1.age = 53 //shortcutでp.Xと書くことも出来る
}
