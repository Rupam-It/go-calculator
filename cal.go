package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}

func main() {
	var a, b int
	fmt.Println("Enter value of a: ")
	fmt.Scan(&a)
	fmt.Println("Enter value of b: ")
	fmt.Scan(&b)

	fmt.Println("for add press 1")
	fmt.Println("for sub press 2")
	fmt.Println("for mul press 3")
	fmt.Println("for div press 4")

	var d, res int
	fmt.Scan(&d)

	if d == 1 {
		res = add(a, b)
	} else if d == 2 {
		res = sub(a, b)
	} else if d == 3 {
		res = mul(a, b)
	} else if d == 4 {
		res = div(a, b)
	} else {
		fmt.Print("your input is in valid ")
	}
	// add some new line
	fmt.Println("this is a added line , has no work with in the repository")
	fmt.Println(res)

}
