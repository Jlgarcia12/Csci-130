package main

import (
	"fmt"
	"html/template"
	//	"log"
	"net/http"
	//	"os"
	"strings"
)

func main() {
	// Add a handler to handle serving static files from a specified directory
	// The reason for using StripPrefix is that you can change the served
	// directory as you please, but keep the reference in HTML the same.
	//http.Handle("/ghostDir/", http.StripPrefix("/ghostDir/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", root)
	http.HandleFunc("/upper", upper)
	// fmt.Println("listening...")
	// err := http.ListenAndServe(GetPort(), nil)
	// if err != nil {
	// log.Fatal("ListenAndServe: ", err)
	// return
	// }
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const upperTemplateHTML = `
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <link rel="stylesheet" href="css/upper.css">
      <title>String Upper Results</title>
    </head>
    <body>
      <h1>String Upper Results</h1>
      <p>That is not the right name <p>
      <pre>{{.}}</pre>
    </body>
  </html>
`

const upperTemplateHTML2 = `
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <link rel="stylesheet" href="css/upper.css">
      <title>String Upper Results</title>
    </head>
    <body>
      <h1>String Upper Results</h1>
      <p>That is the right name: </p>
      <pre>{{.}}</pre>
    </body>
  </html>
`

var upperTemplate = template.Must(template.New("upper").Parse(upperTemplateHTML))
var upperTemplate2 = template.Must(template.New("upper").Parse(upperTemplateHTML2))

func upper(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")
	strUpper := strings.ToUpper(strEntered)

	if strUpper == "JOSE" {
		upperTemplate2.Execute(w, strUpper)
	} else {
		upperTemplate.Execute(w, strUpper)
	}
	// if err != nil {
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <link rel="stylesheet" href="css/upper.css">
        <title>String Upper</title>
      </head>
      <body>
        <h1>String Upper</h1>
        <p>Please enter the best name in the world </p>
        <form action="/upper" method="post" accept-charset="utf-8">
 <input type="text" name="str" placeholder="Type a name" id="str">
 <input type="submit" value="Submit name">
        </form>
      </body>
    </html>
`

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`

// func GetPort() string {
// var port = os.Getenv("PORT")
// // Set a default port if there is nothing in the environment
// if port == "" {
// port = "4747"
// fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
// }
// return ":" + port
// }
