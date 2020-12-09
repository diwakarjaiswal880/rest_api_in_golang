// Golang program to illustrate the
// concept of html/templates
package main

import (
	"html/template"
	"os"
)

// declaring struct
type Student struct {

	// defining struct fields
	Name  string
	Marks int
	Id    string
}

// main function
func main() {

	// defining struct instance
	std1 := Student{"Vani", 94, "20024"}

	// Parsing the required html
	// file in same directory
	t, _ := template.ParseFiles("index.html")

	// standard output to print merged data
	t.Execute(os.Stdout, std1)
}
