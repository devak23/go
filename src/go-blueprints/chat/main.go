package main

import (
  "log"
  "net/http"
  "sync"
  "text/template"
  "path/filepath"
)

type templateHandler struct {
  once sync.Once
  filename string
  tmpl *template.Template
}

// This is the Template Loader function
func (t *templateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  // DO this just once... (this is a singleton). sync.Once guarantees whatever function we pass it
  // will be invoked just once regardless of how many goroutines are calling the function
  t.once.Do(func() {
    tmplFilePath := filepath.Join("template", t.filename) // create the file path
    t.tmpl = template.Must(template.ParseFiles(tmplFilePath)) // parse and load the template
  })
  // compiling the template in the ServeHTTP method also ensures that our code does the work
  // only when its actually needed and not before it
  t.tmpl.Execute(res, nil)
}

func main() {
  http.Handle("/", &templateHandler{filename:"chat.html"})

  // start the web server
  if err := http.ListenAndServe(":8080",nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
