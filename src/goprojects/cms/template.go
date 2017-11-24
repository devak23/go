package cms

import "html/template" //html template library

// template.Must is a utility for template initialization
// template.ParseGlob parses and reads all the templates from a
// directory and loads them into the memory when the application starts

var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title   string
	Content string
}
