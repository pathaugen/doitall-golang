
package main

import "testing"


// ========== ========== ========== ========== ==========
func TestAdd(t *testing.T) {
  t.Log("Adding the numbers 12 and 28... (expected result: 40)")

  result := Add(12,28)

  if result != 40 {
    t.Errorf("Expected result of 40, but it was %d instead.", result)
  }

}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
func TestAddArray(t *testing.T) {
  t.Log("Using an array and SumArray(), Add() all values in a loop... (expected result: 75)")

  _, result := SumArray( []int{ 10, 11, 12, 13, 14, 15 } )

  if result != 75 {
    t.Errorf("Expected result of 75, but it was %d instead.", result)
  }

}
// ========== ========== ========== ========== ==========
