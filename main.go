package main

import (
	"fmt"
	t "gobook/codewars"
)

func main() {
	fmt.Println(t.FirstNonRepeating("a"))
	fmt.Println(t.FirstNonRepeating("stress"))
	fmt.Println(t.FirstNonRepeating("moonmen"))
	fmt.Println(t.FirstNonRepeating(""))
	fmt.Println(t.FirstNonRepeating("abba"))
	fmt.Println(t.FirstNonRepeating("aa"))
	fmt.Println(t.FirstNonRepeating("~><#~><"))
	fmt.Println(t.FirstNonRepeating("sTreSS"))
	fmt.Println(t.FirstNonRepeating("hello world, eh?"))
	fmt.Println(t.FirstNonRepeating("Go hang a salami, I'm a lasagna hog!"))
}
