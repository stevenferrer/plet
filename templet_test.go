package plet

import (
	"os"
	"testing"
)

var (
	contentDir = "examples/tmplt/content/"
	layoutDir  = "examples/tmplt/layout/"
)

func TestNewTemplate(t *testing.T) {
	tmplt := NewBasic()

	if tmplt.ContentDir != "" ||
		tmplt.LayoutDir != "" ||
		tmplt.Ext != "" {

		t.Errorf("New should return an empty template")
	}
}

func TestNewTemplates(t *testing.T) {
	tmplts := NewTemplates()

	tmplt := &Template{
		ContentDir: "path/to/file",
	}

	tmplts.Add(tmplt)
	err := tmplts.Init()
	if err == nil {
		t.Errorf("expecting error to be non-nil")
	}
}

func TestHotReload(t *testing.T) {
	tmplt := NewBasic()
	tmplt.ContentDir = contentDir + "simple"
	tmplt.LayoutDir = layoutDir + "basic"
	tmplt.HotReload = true
	err := tmplt.Init()
	if err != nil {
		t.Errorf("%v", err)
	}

	tmplt.Execute(os.Stdout, nil)
	tmplt.Execute(os.Stdout, nil)
}

func TestTemplatesHotReload(t *testing.T) {
	tmplts := NewTemplates()
	tmplt := New(contentDir+"simple", layoutDir+"basic")
	tmplts.HotReload = true

	err := tmplts.Add(&tmplt)
	if err != nil {
		t.Errorf("error adding template: ", err)
	}

	err = tmplts.Init()
	if err != nil {
		t.Errorf("%v", err)
	}

	tmplt2, err := tmplts.Get("simple")
	if err != nil {
		t.Errorf("Not expecting error\n")
	}
	tmplt2.Execute(os.Stdout, nil)
	tmplt2.Execute(os.Stdout, nil)
}
