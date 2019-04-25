package main 

import "fmt"

/*
The Euclidean Algorithm is used to find the
greatest common divisor of two numbers.
*/

func euclidean_gcd(a int, b int) (int){

	if a < b {
		tmp := a

		a = b
		b = tmp
	}

	r := a % b

	if (r == 0) {
		return b
	}

	return euclidean_gcd(b, r)

}

func main() {

	var a, b int 

	fmt.Println("Enter Two Numbers")
	fmt.Scan(&a, &b)

	gcd := euclidean_gcd(a,b)

	fmt.Println("GCD:", gcd)


}

