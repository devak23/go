package main

import "golearning/src/goprojects/cms"
import "os"

func main() {
	p := &cms.Page{
		Title:   "Hello World",
		Content: "This is the body of the html page",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
