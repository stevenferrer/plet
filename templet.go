package plet

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	//ErrTemplateNotExist returned when not template is not found
	//in the map of templates
	ErrTemplateNotExist = errors.New("template not found")
)

//NewBasic returns a new empty Template
// if you use this, you have to manually add
//the ContentDir (required)
//and LayoutDir and Ext which are both optional
func NewBasic() Template {
	return Template{}
}

//New returns a template using the provided layout and content directory
func New(contentDir, layoutDir string) Template {
	name := filepath.Base(contentDir)
	return Template{
		LayoutDir:  layoutDir,
		ContentDir: contentDir,
		name:       name,
	}
}

//NewTemplates returns a new Templates
//which actually is just a map of Template
func NewTemplates() *Templates {
	var ts Templates
	ts.templates = make(map[string]*Template)
	return &ts
}

//Template wraps the html/template to add functionality
type Template struct {
	//content template directory
	//define the templates here and use
	//them in the layout template
	ContentDir string
	//layout template directory
	//content templates that are used
	//in the layout template must be defined
	//in the content template
	LayoutDir string
	//file name extension of template files
	Ext    string
	layout layout
	//flag to indicate if we're using layout
	useLayout bool
	//flag to indicate if initialization is already performed
	//if false, we are going to return an error
	initialized bool
	//actual template
	t *template.Template

	//set to true to enable hot reload
	HotReload bool

	once sync.Once

	//template name (base directory of template files)
	name string
}

//TODO: Just panic instead of returning an error?

//Execute wraps Template.ExecuteTemplate
func (t *Template) Execute(wr io.Writer, data interface{}) (err error) {
	if t.HotReload {
		if err := t.Init(); err != nil {
			return err
		}
	} else {
		var err error
		t.once.Do(func() {
			err = t.Init()
		})

		if err != nil {
			return err
		}
	}

	if t.useLayout {
		//the layout name will be used as the name of template
		err = t.t.ExecuteTemplate(wr, t.layout.name, data)
	} else {
		err = t.t.Execute(wr, data)
	}

	return
}

// Init performs the initialization
// contentDir: is the directory of content templates
// layoutDir: is the directory of layout templates
// ext: is for file extension (e.g. .html, .txt etc.) .html is used if ext is empty
// Note: if you even forgot to call Init, it will be called when you call Execute
// But it is strongly suggested that you initialize it before everyting else since
// so that it can be verified if the template files exists
func (t *Template) Init() error {

	t.name = filepath.Base(t.ContentDir)

	//ContentDir should be a directory
	aDir, err := isDir(t.ContentDir)
	if err != nil && !aDir {
		return fmt.Errorf("error checking directory %q: %v", t.ContentDir, err)
	}

	if t.Ext == "" {
		t.Ext = ".html"
	}

	if t.LayoutDir != "" {
		//LayoutDir should be a directory
		aDir, err = isDir(t.LayoutDir)
		if err != nil && !aDir {
			return fmt.Errorf("error checking directory %q: %v", t.LayoutDir, err)
		}

		//init layout
		t.layout.init(t.LayoutDir, t.Ext)
		//inicate that we're using layout
		t.useLayout = true
	}

	contentFiles, err := getFilesWithExt(t.ContentDir, t.Ext)
	if err != nil {
		return fmt.Errorf("error getting content files: %v", err)
	}

	//append layout filenames if we're using layouts
	if t.useLayout {
		contentFiles = append(contentFiles, t.layout.fileNames...)
	}

	tmpl := template.Must(template.ParseFiles(contentFiles...))

	t.t = tmpl
	//done initialization
	t.initialized = true

	return nil
}

type layout struct {
	name      string
	fileNames []string
}

//Init gets all the file names for a given
//directory to be used as layout in Content templates
//Note: Directory name should match the layoutName
func (l *layout) init(dir, ext string) error {
	//get the directory name to be
	l.name = filepath.Base(dir)

	fileNames, err := getFilesWithExt(dir, ext)
	if err != nil {
		return fmt.Errorf("error getting files in directory %q: %v", dir, err)
	}

	l.fileNames = fileNames
	return nil
}

//isDir returns true p(path) is a directory
func isDir(p string) (bool, error) {
	if p == "" {
		return false, nil
	}
	var isDir bool
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err == nil && path == p && info.IsDir() {
			isDir = true
		}
		return err
	}

	err := filepath.Walk(p, walkFn)

	return isDir, err
}

//getFilesWithExt gets all files with a specific extention e.g. html
func getFilesWithExt(dir, ext string) ([]string, error) {
	var filenames []string

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(path) == ext {
			filenames = append(filenames, path)
		}
		return err
	}

	if err := filepath.Walk(dir, walkFn); err != nil {
		return nil, err
	}

	return filenames, nil
}
