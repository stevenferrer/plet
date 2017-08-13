package main

import (
	"log"
	"net/http"

	"github.com/steven-ferrer/plet"
)

const (
	port       = ":10123"
	contentDir = "tmplt/content/"
	layoutDir  = "tmplt/layout/"
)

//this example demonstrates rendering template in browser
func main() {
	//new template
	t := plet.New(contentDir+"simple", layoutDir+"basic")

	//initialize template, this will compile the template
	err := t.Init()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//execute
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
