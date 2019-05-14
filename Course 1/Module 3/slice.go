package main

import (
"fmt"
"bufio"
"os"
"strconv"
"strings"
"sort"
)

func main() {
	sli := make([]int, 3)
	var x byte = 'a'
	var i int = 0
	for x != 'X' {
		fmt.Println("Enter a number or 'X' to exit")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
    		n, err := strconv.Atoi(strings.TrimSpace(text))
		if err == nil {
			if i < 3 {
				sli[i] = n
				i = i + 1
			} else {
				sli = append(sli, n)
			}
			sli_copy := append(sli[:0:0], sli...)
			sort.Ints(sli_copy)
			fmt.Println(sli_copy)
		} else {
			if strings.ToUpper(strings.TrimSpace(text)) == "X" {
				x = 'X'
			} else {
				fmt.Println("Incorrect Input")
			}
		}				
	}
}
