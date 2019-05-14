package main

import (
"fmt"
"bufio"
"os"
"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))
	if strings.HasPrefix(text, "i") && strings.HasSuffix(text, "n") && strings.Contains(text, "a"){
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	} 
}