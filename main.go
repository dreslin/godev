package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {

		if number%2 == 0 {
			fmt.Printf("%v is even\n", number)
		} else {
			fmt.Printf("%v is odd\n", number)
		}
	}

	//if number%2 == 0 {
	//	fmt.Println("%v is even", number)
	//} else {
	//	fmt.Println("%v is odd")

	//}

}
