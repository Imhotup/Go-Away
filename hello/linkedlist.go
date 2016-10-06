/**
 * Filename: struct.go
 * Author: Nyah Check
 * Usage: Experiments on Arrays, slices and Maps.
 * Licence: GNU GPL 2016
 */
package main

import (
	"fmt"
)


type st_info struct {

	Name string
	ID string
	Age int
	Dept string
}



func main() {

	//var students [10]st_info
	//students[0] = {"Nyah Check", "FE12A025", 25, "Computer Engineering"}
	
	scores := [4]int{78, 84, 56, 69} //students math scores.
	
	slides := make{[]int,10} //creates a slide structure.
	slides[7] = 43
	 

	for _, value := range scores {
		fmt.Printf("\n%d", value)
	}

}

