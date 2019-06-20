package main

import (
"fmt"
"os"
"sort"
"sync"
)

var wg sync.WaitGroup


func sortArr(a []int, c chan []int) {
	fmt.Println(a)
	sort.Ints(a)
	c <- a
	wg.Done()
}

func main() {
c := make(chan []int)
n := 0
fmt.Println("Enter the number of inputs:")
_, err := fmt.Scanf("%d", &n)
if err != nil {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
a := make([]int, n)
for i := 0; i < n; i++ {
fmt.Println("Enter a new number:")
_, err := fmt.Scanf("%d", &a[i])
if err != nil {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
}
var x int = n/4
if x > 0 {
wg.Add(4)
go sortArr(a[0:x], c)
go sortArr(a[x:2*x], c)
go sortArr(a[2*x:3*x], c)
go sortArr(a[3*x:n], c)
}
newArr := []int{}
for i := 0; i < 4; i++ {
	for _, v := range <-c {
			newArr = append(newArr, v)
		}
}
wg.Wait()
sort.Ints(newArr)
fmt.Println("Sorted Array: ", newArr)
}
