package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/steven-ferrer/plet"
)

const (
	port       = ":10123"
	contentDir = "tmplt/content/"
	layoutDir  = "tmplt/layout/"
)

// This example demonstrates many of the capabilities of
// the plet package for managing templates

func main() {
	tmplts := makeTemplates(contentDir, layoutDir)

	http.HandleFunc("/simple", func(w http.ResponseWriter, r *http.Request) {
		//use the base directory name of 
		//content dir to get the content template
		t, err := tmplts.Get("simple")
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/simple2", func(w http.ResponseWriter, r *http.Request) {
		//use the base directory name of 
		//content dir to get the content template
		t, err := tmplts.Get("simple2")
		if err != nil {
			log.Fatal(err)
		}

		//executing templates is very familiar
		err = t.Execute(w, struct{ Name string }{"John Doe"})
		if err != nil {
			log.Fatal(err)
		}
	})

	//it is also possible to use it without a layout template
	http.HandleFunc("/nolayout", func(w http.ResponseWriter, r *http.Request) {
		//use the base directory name of 
		//content dir to get the content template
		t, err := tmplts.Get("nolayout")
		if err != nil {
			log.Fatal(err)
		}

		//executing templates is very familiar
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

//makeTemplets is used for loading template filenames
func makeTemplates(contentsDir, layoutsDir string) *plet.Templates {
	//contentsDir and layoutsDir should be relative
	//to the executable program

	//we're storing the template config into a json file
	b, err := ioutil.ReadFile("templates.json")
	if err != nil {
		log.Fatalf("error reading templates.json: %v", err)
	}

	var contents tmpltCfg
	err = json.Unmarshal(b, &contents)
	if err != nil {
		log.Fatalf("error unmarshaling templates.json: %v", err)
	}

	tmplts := plet.NewTemplates()

	//set to true on development, by default,
	//once templates are initialized, they will
	//not be re-compiled until you restart the server
	//setting HotReload to true will re-compile the templates
	//everytime Execute is called
	tmplts.HotReload = true

	for _, c := range contents.ContentTmplts {
		t := plet.New(c.ContentDir, c.LayoutDir)
		//templates.Add uses the base directory of
		//the ContentDir as map key
		tmplts.Add(&t)
	}

	//init all templates to make sure everything is fine
	err = tmplts.Init()
	if err != nil {
		log.Fatal(err)
	}

	return tmplts
}

type tmpltCfg struct {
	ContentTmplts []contentTmplts `json:"contents"`
}

type contentTmplts struct {
	//Name       string `json:"name"`
	ContentDir string `json:"contentdir"`
	LayoutDir  string `json:"layoutdir"`
}
