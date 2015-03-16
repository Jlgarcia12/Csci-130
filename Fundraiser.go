package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Honorific}} {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is unfortunate that you could not make it to the fundraiser.{{end}}
{{if .Donated}}Thank you for the generous donation!{{end}}
Upcoming events your invited to are the{{range.Event}}, {{.}}{{end}} fundraiser's.
Best wishes,
Garcia
`

	// Prepare some data to insert into the template.
	type Client struct {
		Honorific, Name   string
		Attended, Donated bool
		Event             []string
	}

	//type Events struct {
	//	upcomingEvents []string
	//}

	var clients = []Client{
		{"Mr.", "Plumber", true, true, []string{"Chicken", "Goose", "Duck"}},
		{"Mrs.", "Garcia", false, false, []string{"Soup", "Bread", "Meat"}},
		{"Ms.", "Lundenberg", false, true, []string{"Poor", "Penniless", "Broke"}},
	}

	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	//STEP 3
	// Execute the template for each recipient.
	for _, r := range clients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

// FROM : http://golang.org/pkg/text/template/#example_Template
