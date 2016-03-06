package main

import "fmt"

func main() {

	s := 1
	for i := 0; i < 10; i++ {
		dbLpush("ksel_list", "lesk"+string(s))
		s = s + 1
	}

	for i := 0; i < 20; i++ {
		fmt.Println(dbLpop("ksel_list"))
	}
}
