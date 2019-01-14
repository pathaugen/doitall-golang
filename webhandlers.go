
package main

import (
	"fmt"

	"net/http"

	//"strconv" // https://golang.org/pkg/strconv/
	//"strings"

	//"bufio"
	//"os"
	//"os/exec"

	//"log"
)


var htmlTemplate = `
  <html>
    <head>
      <title>PatHaugen's doitall-golang</title>
    </head>
    <body style="padding:2%;">
      <h1>doitall-golang</h1>
      <div>
        <span><a href="/">Home</a></span>
        <span><a href="/api">Api</a></span>
      </div>
    </body>
  </html>
`


func handlerApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf( w, "This is /api of the Webserver" )
}
