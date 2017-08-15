package main

/**
 * The purpose of the program is to download the dependencies
 * required by the programs. It works very similar to npm.
 * The possible commands with gpm are:
 * gpm install
 * gpm install <dependency_name> --save
 */

import (
  "encoding/json"
  "fmt"
  "flag"
  "io/ioutil"
  "os"
  "os/exec"
  "sync"
)

type Package struct {
  Name            string `json: "name"`
  Version         string `json: "version"`
  Description     string `json: "description"`
  Main            string `json: "main"`
  Dependencies    map[string]string `json: "dependencies"`
  DevDependencies map[string]string  `json: "devDependencies"`
  Scripts         map[string]string `json: "scripts"`
  Repository      map[string]string  `json: "repository"`
  Keywords        []string `json: "keywords"`
  Author          string  `json: author`
  License         string  `json: license`
  Bugs            map[string]string    `json: bugs`
  Homepage        map[string]string    `json: homepage`
}

func isEmpty(element string) bool {
  return len(element) == 0
}

func main() {
  var wg sync.WaitGroup

  // the flag.String returns a pointer not the value
  flgSave := flag.String ("save", "", "instructs gpm to write the dependency into the gopackage.json file")
  flgInstall := flag.String ("install", "", "instructs go to install a dependency")
  fmt.Printf("flgInstall = %s\n", *flgInstall)
  fmt.Printf("flgSave = %s\n", *flgSave)

  // parse the command line flags
  flag.Parse()

  if len(os.Args) == 1 || isEmpty(*flgInstall)  {
    fmt.Printf("Usage:\tgpm install OR\n\tgpm install <dependency-name> -save\n")
    os.Exit(1)
  }

  // check for the existence of gopackage.json
  mypackage := readPackage()

  //TODO: check if there is another dependency specified with --save flag

  // if no other flags are passed, simply install all dependencies
  if isEmpty(*flgSave) {
    installAllDependencies(mypackage, &wg)
  }

  // pause for all the goroutines to end
  wg.Wait()
}

/**
 * This function reads file and transforms it into the Package object
 */
func readPackage() Package {
  raw, err := ioutil.ReadFile("./gopackage.json")
  if err != nil {
    fmt.Printf("Problem reading gopackage.json: %s\n", err)
    os.Exit(1)
  }

  var mypackage Package
  json.Unmarshal(raw, &mypackage)

  return mypackage
}

/**
 * This function will spawn several goroutines and will download all
 * dependencies described in the package
 */
func installAllDependencies(mypackage Package, wg *sync.WaitGroup) {

  // TODO: need to check if the dependency is already downloaded
  // TODO: what if one of the dependency is removed from the gopackage.json

  // read the gopackage.json's dependencies section
  for k,_ := range mypackage.Dependencies {
    wg.Add(1)
    fmt.Printf("Dependency = %s\n", k)
    go func(packageName string) {
      fetch(packageName)
      wg.Done()
    }(k)
  }
}

/**
 * This function runs the "go get <dependency>" command to download
 * the dependency into the current folder
 */
func fetch(packageName string) error {
  fmt.Printf("downloading the package %s...\n", packageName)
  out, err := exec.Command("go", "get", packageName).Output()
  if err != nil {
    return err
  }
  fmt.Printf("output: %s\n", out)
  return err
}
