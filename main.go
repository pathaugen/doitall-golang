
package main

import (
	"fmt"

	"strconv" // https://golang.org/pkg/strconv/
)


var appinfo = `
  =======================================
       _           _  _             _  _
      | |         (_)| |           | || |
    __| |  ___     _ | |_     __ _ | || |
   / _`+"`"+` | / _ \   | || __|   / _`+"`"+` || || |
  | (_| || (_) |  | || |_   | (_| || || |
   \__,_| \___/   |_| \__|   \__,_||_||_|
                             ..in Golang!
  =======================================
`


// TODO: Read from CHANGELOG.md and output the text
var changelog = `
CHANGELOG:

Version 0.0.1 BETA
 - Splash ASCII
 - Basic functions
 - Reading flags from command line
`


var documentation = `
$ doitall-golang
  Runs the application.

$ doitall-golang -changelog
  View the changelog.

---------- Math Operations ----------

  Addition:
  $ doitall-golang -add X Y
    Adds the variables X and Y

  Subtraction:
  $ doitall-golang -sub X Y
    Adds the variables X and Y

  Multiplication:
  $ doitall-golang -mult X Y
    Adds the variables X and Y

  Division:
  $ doitall-golang -div X Y
    Adds the variables X and Y

---------- API Operations ----------

  Define API Endpoint:
  $ doitall-golang -api X
    Specifies X as the API endpoint to utilize for API operations.

`


func add(x int, y int) int {
	return x + y
}


func main() {

	breakline := ("\n\n" + "========== ========== ========== ========== ==========" + "\n\n")
	breakspace := ("\n\n")

	fmt.Print(breakspace)

	fmt.Print("Do it all in Golang!")
	fmt.Print(breakspace)
	fmt.Print("https://github.com/pathaugen/doitall-golang")

	fmt.Print(breakline)

	// ========== ========== ========== ========== ==========
	// Simple Addition
	fmt.Print("Simple Addition:" + breakspace)
	//fmt.Print("Calling function add with 42, 13: " + add(42, 13))
	fmt.Print(add(42, 13))
	// ========== ========== ========== ========== ==========

	fmt.Print(breakline)

	// ========== ========== ========== ========== ==========
	// Simple Loop
	fmt.Print("Simple Loop:" + breakspace)
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Print("interval: " + strconv.Itoa(i) + "\n") // Itoa is shorthand for FormatInt(int64(i), 10)
		sum += i
		fmt.Print("sum: " + strconv.Itoa(sum) + "\n")
	}
	fmt.Print("Final sum: " + strconv.Itoa(sum))
	// ========== ========== ========== ========== ==========

	fmt.Print(breakline)

	// ========== ========== ========== ========== ==========
	// Take Command Line Arguments
	fmt.Print("Take Command Line Arguments:" + breakspace)
	// ========== ========== ========== ========== ==========

	fmt.Print(breakline)

	// ========== ========== ========== ========== ==========
	// Load JSON into Array
	fmt.Print("Load JSON into Array:" + breakspace)
	// ========== ========== ========== ========== ==========

	fmt.Print(breakline)

}
