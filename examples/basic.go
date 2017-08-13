package main

import (
	"log"
	"os"

	"github.com/steven-ferrer/plet"
)

const (
	contentDir = "tmplt/content/"
	layoutDir  = "tmplt/layout/"
)

//very simple example of using plet
func main() {
	//make a new template
	t := plet.New(contentDir+"simple", layoutDir+"basic")

	//initialize template, this will compile the template
	err := t.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
