package main

import (
"fmt"
"bufio"
"os"
"strings"
)

type person struct {
	fname string 
	lname string
	
}

func main() {
	sli := make([]person, 0)
	fmt.Println("Enter path of the file:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	file, err := os.Open(strings.TrimSpace(name))
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		sli = append(sli, person{fname: s[0], lname: s[1]})
	}

	for i := 0; i <= len(sli[1:]); i++ {
		fmt.Println("Struct No.", i + 1)
		fmt.Println("FName: ", sli[i].fname)
		fmt.Println("LName: ", sli[i].lname, "\n")
	}
}
