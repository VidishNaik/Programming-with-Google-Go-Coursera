package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Enter a floating point number:")
	fmt.Scan(&x)
	var y = int(x)
	fmt.Println(y)
}
