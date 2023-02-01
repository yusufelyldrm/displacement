package main

import "fmt"

// GenDisplaceFn returns a function that computes displacement as a function of time, given acceleration,
// initial velocity, and initial displacement.
// The function returned by GenDisplaceFn takes one argument (time), computes the displacement, and returns the value as a float64.
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
	return fn
}

func main() {
	var a, v0, s0, t float64

	fmt.Println("Enter acceleration:")
	fmt.Scan(&a)

	fmt.Println("Enter initial velocity:")
	fmt.Scan(&v0)

	fmt.Println("Enter initial displacement:")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println("Enter time:")
	fmt.Scan(&t)

	fmt.Println("Your location: ", fn(t))

}
