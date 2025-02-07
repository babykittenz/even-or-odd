package main

import "fmt"

type numbers []int

func (n numbers) evenOrOdd() []string {
	evenOrOddHolder := []string{}
	for _, number := range n {
		if number%2 == 0 {
			evenOrOddHolder = append(evenOrOddHolder, "even")
			fmt.Println("even,", number)
		} else {
			evenOrOddHolder = append(evenOrOddHolder, "odd")
			fmt.Println("odd,", number)
		}
	}
	fmt.Println(evenOrOddHolder)
	return evenOrOddHolder
}
