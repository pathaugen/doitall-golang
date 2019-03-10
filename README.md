
Do it all! in Golang
====================
[![CircleCI Build Status](https://circleci.com/gh/pathaugen/doitall-golang.svg?style=svg)](https://circleci.com/gh/pathaugen/doitall-golang)
[![TravisCI Build Status](https://travis-ci.com/pathaugen/doitall-golang.svg?branch=master)](https://travis-ci.com/pathaugen/doitall-golang)
[![Build status](https://ci.appveyor.com/api/projects/status/2hm4g4of8d2k5of1/branch/master?svg=true)](https://ci.appveyor.com/project/PatrickHaugen/doitall-golang/branch/master)

[![Coverage Status](https://coveralls.io/repos/github/pathaugen/doitall-golang/badge.svg?branch=master)](https://coveralls.io/github/pathaugen/doitall-golang?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/4fa3e21b6359a2da09ce/maintainability)](https://codeclimate.com/github/pathaugen/doitall-golang/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/4fa3e21b6359a2da09ce/test_coverage)](https://codeclimate.com/github/pathaugen/doitall-golang/test_coverage)

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
     * https://git-scm.com/
   * Test Git Installation:
     * > git --version
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
