package main

import (
"fmt"
"bufio"
"os"
"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}

type Snake struct {}

type Bird struct {}

func (c Cow) Eat() {
	fmt.Println("Grass")
}

func (c Cow) Move() {
	fmt.Println("Walk")
}

func (c Cow) Speak() {
	fmt.Println("Walk")
}

func (s Snake) Eat() {
	fmt.Println("Mice")
}

func (s Snake) Move() {
	fmt.Println("Slither")
}

func (s Snake) Speak() {
	fmt.Println("Hiss")
}

func (b Bird) Eat() {
	fmt.Println("Worms")
}

func (b Bird) Move() {
	fmt.Println("Fly")
}

func (b Bird) Speak() {
	fmt.Println("Peep")
}

func newAnimal(split []string){
	if split[2] == "cow" {
		animals[split[2]] = Cow{}
	} else if split[2] == "snake" {
		animals[split[2]] = Snake{}
	} else if split[2] == "bird" {
		animals[split[2]] = Bird{}
	}
	fmt.Println("Created it")
}

func query(split []string){
	if animal, ok := animals[split[1]]; ok { 
		if split[2] == "eat" {
			animal.Eat()
		} else if split[2] == "move" {
			animal.Move()
		} else if split[2] == "speak" {
			animal.Speak()
		}
	}
}


var animals = make(map[string]Animal)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		s, _ := reader.ReadString('\n')
		split := strings.Split(strings.ToLower(strings.TrimSpace(s)), " ")
		if len(split) != 3 {
			fmt.Println("Invalid length of input detected. Please re-enter the input.")
			continue
		}
		if split[0] == "newanimal" {
			newAnimal(split)
		} else if split[0] == "query" {
			query(split)
		}
	}
}
