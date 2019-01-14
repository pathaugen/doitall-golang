
package main

import (
	"fmt"

	"net/http"

	//"strconv" // https://golang.org/pkg/strconv/
	"strings"

	"bufio"
	"os"
	"os/exec"

	"time"
	//"log"
)



// NOTES:
// Golang Slices are basically Arrays
// Golang Maps are basically Dicts

// TODO:
// - reorder array alphabetically
// - pull json and parse out, putting into struct first and utilizing there


// Pull changelog.md and display

// Pull readme.md as documenation and display


// ========== START: Golang Console Colors ========== ========== ========== ==========
// Example: fmt.Print( cRed + "HelloWorld" + cClr )
var cClr				= "\u001b[0m"

var cBold				= "\u001b[1m"

var cBlack			= "\u001b[30m"
var cRed				= "\u001b[31m"
var cGreen			= "\u001b[32m"
var cYellow			= "\u001b[33m"
var cBlue				= "\u001b[34m"
var cMagenta		= "\u001b[35m"
var cCyan				= "\u001b[36m"
var cWhite			= "\u001b[37m"

var cBlackBG		= "\u001b[40m"
var cRedBG			= "\u001b[41m"
var cGreenBG		= "\u001b[42m"
var cYellowBG		= "\u001b[43m"
var cBlueBG			= "\u001b[44m"
var cMagentaBG	= "\u001b[45m"
var cCyanBG			= "\u001b[46m"
var cWhiteBG		= "\u001b[47m"
// ========== END: Golang Console Colors ========== ========== ========== ==========


var appinfo = `
  ` + cBlue + `=======================================` + cBold + cCyan + `
       _           _  _             _  _
      | |         (_)| |           | || |
    __| |  ___     _ | |_     __ _ | || |
   / _`+"`"+` | / _ \   | || __|   / _`+"`"+` || || |
  | (_| || (_) |  | || |_   | (_| || || |
   \__,_| \___/   |_| \__|   \__,_||_||_|` + cClr + `
                             ` + cCyan + `..in Golang!` + cClr + `
  ` + cBlue + `=======================================` + cClr + `
`


// > go build&doitall-golang.exe
func main() {


	// ========== ========== ========== ========== ==========
	// Declare and Server in one go
	webserverRunning := false
	http.HandleFunc( "/", func(w http.ResponseWriter, r *http.Request ) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		//fmt.Fprintf( w, "This is / of the Webserver" )
		fmt.Fprintf( w, htmlTemplate )
	})
	http.HandleFunc("/api", handlerApi)
	// ========== ========== ========== ========== ==========


	// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go#22896706
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	breakspace := ("\n")
	//breakline := ( breakspace + cBlue + "  ========== ========== ========== ========== ==========" + cClr + breakspace )
	breakline := ( breakspace + cBlue + "  =======================================" + cClr + breakspace )

	// Start Reading Input
	scanner := bufio.NewScanner(os.Stdin)
	var inputText string

	for ( inputText != "q" ) {  // break the loop if inputText == "q"

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// ========== ========== ========== ========== ==========
		// Take Command Line Arguments
		// fmt.Print("Take Command Line Arguments:" + breakspace)
		// ========== ========== ========== ========== ==========

		fmt.Print(
			appinfo + breakspace +
			"  Do it all in Golang!" + breakspace +
			"  https://github.com/pathaugen/doitall-golang" + breakspace )

		//var b []byte = make([]byte, 1)
		//for {
			//os.Stdin.Read(b)
			//fmt.Println("I got the byte", b, "("+string(b)+")")
		//}

		if ( inputText == "s" ) {

			fmt.Print( breakline )

			// ========== START: Golang Internal Webserver ========== ========== ========== ==========
			// Channels: https://tour.golang.org/concurrency/2
			fmt.Print(
				breakspace + cBold + cCyan + "  Start/Stop Internal Webserver:" + cClr + breakspace )

			fmt.Print( breakspace )

			//webserverChannel := make(chan bool)
			//webserverChannel <- false

			fmt.Print( cBold + cGreen + "  Starting Internal Webserver on localhost port :9090" + cClr )
			go func() {
				//fmt.Print( "Starting Webserver on localhost port :9090" )

				//webserverChannel <- true

				http.ListenAndServe(":9090", nil)
				//time.Sleep(time.Second * 3)

				//webserverChannel <- false
			}()
			webserverRunning = true

			// log.Fatal( http.ListenAndServe(":9090", nil) )
			//err := http.ListenAndServe(":9090", nil) // set listen port
			//if err != nil {
				//log.Fatal("ListenAndServe: ", err)
			//}

			//fmt.Print( <-webserverChannel )
			//for i := 0; i < len(webserverChannel); i++ {
				//fmt.Print( <-webserverChannel )
			//}

			// ========== END: Golang Internal Webserver ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "t" ) {

			fmt.Print( breakline )

			// ========== ========== ========== ========== ==========
			// Run Tests
			fmt.Print(
				breakspace + cBold + cCyan + "  Golang Tests:" + cClr + breakspace )

			fmt.Print( breakspace )

			testOutput, _ := exec.Command("cmd", "/c", "go test -v").Output()
			//cmd.Stdout = os.Stdout
			//cmd.Run()

			// Add Colors to testOutput
			testOutputString := string(testOutput)
			testOutputString = strings.Replace( testOutputString, "RUN", cCyan + "RUN" + cClr, -1 )
			testOutputString = strings.Replace( testOutputString, "PASS", cGreen + "PASS" + cClr, -1 )
			testOutputString = strings.Replace( testOutputString, "FAIL", cRed + "FAIL" + cClr, -1 )

			fmt.Print( testOutputString )

			// ========== ========== ========== ========== ==========

			//fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "0" ) {

			fmt.Print( breakline )

			// ========== ========== ========== ========== ==========
			// Color Test

			fmt.Print(
				breakspace + cBold + cCyan + "  Golang Console Color Testing:" + cClr + breakspace +

				breakspace + "  Basic Text:" + cClr + cBold + "    Bold Text:" + cClr + "     Basic Background:" + cClr +
				breakspace +
				breakspace + cBlack + "  Black" + cClr + "          " + cBold + cBlack + "Black" + cClr + "          " + cBlackBG + "Black" + cClr +
				breakspace + cRed + "  Red" + cClr + "            " + cBold + cRed + "Red" + cClr + "            " + cRedBG + "Red" + cClr +
				breakspace + cGreen + "  Green" + cClr + "          " + cBold + cGreen + "Green" + cClr + "          " + cGreenBG + "Green" + cClr +
				breakspace + cYellow + "  Yellow" + cClr + "         " + cBold + cYellow + "Yellow" + cClr + "         " + cYellowBG + "Yellow" + cClr +
				breakspace + cBlue + "  Blue" + cClr + "           " + cBold + cBlue + "Blue" + cClr + "           " + cBlueBG + "Blue" + cClr +
				breakspace + cMagenta + "  Magenta" + cClr + "        " + cBold + cMagenta + "Magenta" + cClr + "        " + cMagentaBG + "Magenta" + cClr +
				breakspace + cCyan + "  Cyan" + cClr + "           " + cBold + cCyan + "Cyan" + cClr + "           " + cCyanBG + "Cyan" + cClr +
				breakspace + cWhite + "  White" + cClr + "          " + cBold + cWhite + "White" + cClr + "          " + cWhiteBG + "White" + cClr )
			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "1" ) {

			fmt.Print( breakline )

			// ========== ========== ========== ========== ==========
			// Simple Addition: Add(x,y)
			fmt.Print(
				breakspace + cBold + cCyan + "  Simple Addition: Add(x,y)" + cClr + breakspace )

			fmt.Print( breakspace )

			//fmt.Print("Calling function add with 42, 13: " + add(42, 13))

			fmt.Print( Add(42, 13) )
			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "2" ) {

			fmt.Print( breakline )

			// ========== ========== ========== ========== ==========
			// Add Array Values in Loop: SumArray( array )
			fmt.Print(
				breakspace + cBold + cCyan + "  Add Array Values in Loop: SumArray( array )" + cClr + breakspace )

			fmt.Print( breakspace )

			// Define an Array
			inputArray := []int{ 10, 11, 12, 13, 14, 15 }
			funcOutput, _ := SumArray( inputArray )
			fmt.Print( funcOutput )

			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "3" ) {

			fmt.Print( breakline )

			// ========== ========== ========== ========== ==========
			// Load JSON into Array
			fmt.Print(
				breakspace + cBold + cCyan + "  Load JSON into Array:" + cClr + breakspace )
			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "4" ) {

			// ========== ========== ========== ========== ==========
			// Reorder Array Alphabetically
			fmt.Print(
				breakspace + cBold + cCyan + "  Reorder Array Alphabetically:" + cClr + breakspace )
			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "5" ) {

			// ========== ========== ========== ========== ==========
			fmt.Print(
				breakspace + cBold + cCyan + "  Goroutines Communicating via Channels" + cClr + breakspace )

			goroutinecheck := make(chan int)
			go func() {
				time.Sleep(time.Second * 3)
				goroutinecheck <- 1
			}()
			go func() {
				time.Sleep(time.Second * 2)
				goroutinecheck <- 2
			}()
			go func() {
				time.Sleep(time.Second * 1)
				goroutinecheck <- 3
			}()
			for i := 0; i < 3; i++ {
				fmt.Print( <-goroutinecheck )
			}

			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else if ( inputText == "6" ) {

			// ========== ========== ========== ========== ==========
			fmt.Print(
				breakspace + cBold + cCyan + "  OPTION 6:" + cClr + breakspace )
			// ========== ========== ========== ========== ==========

			fmt.Print( breakspace )
			fmt.Print( breakline )
			fmt.Print( breakspace )
			fmt.Print( "  [" + cGreen + "enter" + cClr + "] " + cCyan + "Return To Menu" + cClr + breakspace + "      [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()

		} else {

			// ========== ========== ========== ========== ==========
			fmt.Print(breakline)
			fmt.Print( breakspace + cBold + cCyan + "  Application Menu:" + cClr + breakspace + breakspace )

			// Webserver Status
			webserverStatus := ``
			if ( webserverRunning ) {
				webserverStatus = cBold + cGreen + "RUNNING" + cClr
			} else {
				webserverStatus = cBold + cRed + "STOPPED" + cClr
			}

			fmt.Print(
				"  [" + cGreen + "0" + cClr + "] " + cCyan + "Golang Console Color Testing" + cClr + breakspace +
				"  [" + cGreen + "1" + cClr + "] " + cCyan + "Simple Addition: Add(x,y)" + cClr + breakspace +
				"  [" + cGreen + "2" + cClr + "] " + cCyan + "Add Array Values in Loop: SumArray( array )	" + cClr + breakspace +
				"  [" + cGreen + "3" + cClr + "] " + cCyan + "Load JSON into Array" + cClr + breakspace +
				"  [" + cGreen + "4" + cClr + "] " + cCyan + "Reorder Array Alphabetically" + cClr + breakspace +
				"  [" + cGreen + "5" + cClr + "] " + cCyan + "Goroutines Communicating via Channels" + cClr + breakspace +
				"  [" + cGreen + "6" + cClr + "] " + cCyan + "Option 6" + cClr + breakspace +

				"  [" + cRed + "s" + cClr + "] " + cCyan + "Start/Stop Internal Webserver " + cClr + "[" + webserverStatus + "]" + cClr + breakspace +
				"  [" + cRed + "t" + cClr + "] " + cCyan + "Run Golang Tests" + cClr + "              " + "\\ http://localhost:9090" + breakspace + // Unicode: ↪
				"  [" + cRed + "q" + cClr + "] " + cCyan + "Quit Application" + cClr + breakspace )


			fmt.Print(breakline)
			fmt.Print( breakspace )
			fmt.Print( cBold + cCyan + "  Enter a Numeric Selection: " + cClr + breakspace + cYellow + "  > " + cClr )
			scanner.Scan()
			inputText = scanner.Text()


			if inputText != "q" {
				//fmt.Println("Your text was: ", inputText)
			}

			// selectedOption := 0

			// Array of options, if i = selectedOption, draw carrot in front ">"

			// fmt.Print(breakline)
			// ========== ========== ========== ========== ==========

		}
	}
}
