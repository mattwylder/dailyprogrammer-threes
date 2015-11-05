package main

import "fmt"
import "math/rand"

func main() {
    num := rand.Int()
    for num > 1 {
	if num % 3 == 0 {
	    num = num /3
	} else {
		num = num + 1
	}
	fmt.Println(num)
    }
}
