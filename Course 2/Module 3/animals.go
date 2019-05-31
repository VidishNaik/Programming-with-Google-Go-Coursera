package main

import (
"fmt"
"bufio"
"os"
"strings"
"reflect"
)

type Animal struct {
	eat, move, speak string
}

func (a Animal) Eat() string {
	return a.eat
}

func (a Animal) Move() string {
	return a.move
}

func (a Animal) Speak() string {
	return a.speak
}

func checkInput(split []string) bool{
	if split[0] == "cow" || split[0] == "bird" || split[0] == "snake"{
		if split[1] == "eat" || split[1] == "move" || split[1] == "speak" {
			return true
		} else {
			fmt.Println("Action not present.")
		}
	} else {
		fmt.Println("Animal not present.")
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cow := Animal{"Grass", "Walk", "Moo"}
	_ = cow	
	bird := Animal{"Worm", "Fly", "Peep"}
	snake := Animal{"Mice", "Slither", "Hiss"}
	for {
		fmt.Print("> ")
		s, _ := reader.ReadString('\n')
		split := strings.Split(strings.ToLower(strings.TrimSpace(s)), " ")
		if len(split) != 2 {
			fmt.Println("Invalid length of input detected. Please re-enter the input.")
			continue
		}
		if checkInput(split) {
			var obj Animal
			if split[0] == "cow" {
				obj = cow				
			} else if split[0] == "bird" {
				obj = bird		
			} else if split[0] == "snake" {
				obj = snake			
			}
			fmt.Println(reflect.ValueOf(obj).MethodByName(strings.Title(split[1])).Call(nil))
		}
	}
}
