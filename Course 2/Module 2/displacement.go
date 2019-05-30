package main

import (
"fmt"
)

func main() {
var a, v0, s0, t float64
fmt.Println("Enter the acceleration:")
fmt.Scanf("%f", &a)
fmt.Println("Enter the initial velocity:")
fmt.Scanf("%f", &v0)
fmt.Println("Enter the initial displacement:")
fmt.Scanf("%f", &s0)
fn := GenDisplaceFn(a, v0, s0)
fmt.Println("Enter the time:")
fmt.Scanf("%f", &t)
fmt.Println(fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func (t float64) float64 {
	return func (t float64) float64 {
		return ((0.5 * a * t * t) + (v0 * t) + s0)	
	}
}
