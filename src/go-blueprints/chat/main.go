package main

import (
  "log"
  "net/http"
  "sync"
  "text/template"
  "path/filepath"
  "flag"
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
  t.tmpl.Execute(res, req)
}

func main() {
  var addr = flag.String("addr", ":8080","The address of the application")
  flag.Parse()

  r := newRoom()
  http.Handle("/", &templateHandler{filename:"chat.html"})
  http.Handle("/room", r)

  // get the room going
  go r.run()

  // start the web server
  log.Println("Starting web server on", *addr)
  if err := http.ListenAndServe(*addr,nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
