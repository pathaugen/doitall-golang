
Do it all! in Golang
====================

Description
-----------

Simple project framework to provide code and a baseline to do just about anything in Go.

Quick Start
------------

1. Machine Setup
2. Compile and Run the Application
3. Operations Available:
   * Math Operations:
     * Addition
     * Subtraction
     * Multiplication
     * Division
   * API Operations:
     * Define API Endpoint

---------- Machine Setup ----------

1. Install Go:
   * https://golang.org/dl/
   * Recommend C:\Go on Windows
   * Test Go Installation:
     * > go version
2. Setup Go:
   * Create File Structure:
     * C:\go-work\bin
     * C:\go-work\pkg
     * C:\go-work\src
   * Check Environment Variables:
     * GOROOT: C:\Go
     * GOPATH: C:\go-work
3. Setup Git:
   * Install Git:
   * x
x. Test Code:
   * > go test -v
   * xxx
x. Check Documentation:
   * > godoc -http=:6060
   * http://localhost:6060/pkg/
     * All code in $GOPATH (which includes stdlib) appears in web interface.

---------- Compile and Run the Application ----------

$ go build
  Compiles the application.

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

Features
--------

(developing)

External Libraries
------------------

(developing)
