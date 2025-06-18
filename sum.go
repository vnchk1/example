package main

import "fmt"

func sum(n,m int) float64 {
	return (n + m) / 3
}
func main() {
	fmt.Println(sum(4, 4))
	fmt.Println("bugFix")
}