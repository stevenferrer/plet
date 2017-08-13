## Plet

Plet is a templet nesting package for Go. Plet wraps html/template and provides helper functions to make it easy to manage and use html/template.

## Terminologies

__Layout__: layout templates serves as base for content templates. 
__Content__: content templates is where you define the sections of your templates.

## Basic Usage

[tmplt/layout/basic/basic.html](https://github.com/steven-ferrer/plet/blob/master/examples/tmplt/layout/basic/basic.html): layout template

	{{ define "basic" }}
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			{{ template "head" . }}
		</head>
		<body>
			{{ template "content" . }}
			</body>
	</html>
	{{ end }}

[tmplt/content/simple/content.html](https://github.com/steven-ferrer/plet/blob/master/examples/tmplt/content/simple/content.html): content template

	{{ define "head" }}
	<title>Plet | Template Management Package</title>
	{{ end }}

	{{ define "content" }}
	<h1>Hello World!</h1>
	{{ end }}


[basic.go](https://github.com/steven-ferrer/plet/blob/master/examples/basic.go):

	import (
		"log"
		"os"

		"github.com/steven-ferrer/plet"
	)

	const (
		contentDir = "tmplt/content/"
		layoutDir  = "tmplt/layout/"
	)

	//very simple example of using plet package
	func main() {
		//new template
		t := plet.New(layoutDir+"basic", contentDir+"simple")
		
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

## More Examples

For more examples please see [examples](https://github.com/steven-ferrer/plet/tree/master/examples) folder.

## Issues and Question

If you found a bug or have a question, please feel free to file an issue.

## Contributing

Please feel free to contribute by openning an PR.
