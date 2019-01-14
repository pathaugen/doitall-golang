
package main

import (
	//"fmt"

	"strconv" // https://golang.org/pkg/strconv/
)


// ========== START: Add(x,y) Addition ========== ========== ========== ==========
func Add(x int, y int) int {
	return x + y
}
// ========== END: Add(x,y) Addition ========== ========== ========== ==========


// Subtraction

// Multiplication

// Division


// ========== START: SumArray( array ) ========== ========== ========== ==========
// Take an array input and add all values together
// Can break array in half and utilize goroutines and a channel to split the work: https://tour.golang.org/concurrency/2
func SumArray( inputArray []int ) (string, int) {
  outputString := ``
  resultInt := 0


  // ==========
  //for i := 0; i < 10; i++ {
    //outputString += "  interval: " + strconv.Itoa(i) + "\n" // Itoa is shorthand for FormatInt(int64(i), 10)
    // //sum += i
    //resultInt = Add( resultInt, i)
    //outputString += "    sum: " + strconv.Itoa(resultInt) + "\n"
  //}
  // ==========

  // ==========
  for i := 0; i < len(inputArray); i++ {
    outputString += "  interval: " + strconv.Itoa(i) + "\n" // Itoa is shorthand for FormatInt(int64(i), 10)
    resultInt = Add( resultInt, inputArray[i] )
    outputString += "    sum: " + strconv.Itoa(resultInt) + "\n"
  }
  // ==========

  outputString += "  Final sum: " + strconv.Itoa(resultInt)

	return outputString, resultInt
}
// ========== END: SumArray( array ) ========== ========== ========== ==========
