package main

import "fmt"

func main() {
	fmt.Println(split(17))

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	//A return statement without arguments returns the named return values.
	//This is known as a "naked" return.
	return
}
