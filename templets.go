package plet

//Templates is a map of plet.Template
type Templates struct {
	//set to true if you want all the
	//tempaltes contained in here to enable hot reload
	HotReload bool

	templates map[string]*Template
}

//Add adds a template to the template map
//if template is previously added, it will be overwritten
func (tmplts *Templates) Add(name string, tmplt *Template) {
	//match the hot reload to contained templates
	//if hot reload is enabled
	if tmplts.HotReload {
		tmplt.HotReload = true
	}
	tmplts.templates[name] = tmplt
}

//Get returns the template given a name
func (tmplts *Templates) Get(name string) (*Template, error) {
	t, ok := tmplts.templates[name]
	if !ok {
		return &Template{}, ErrTemplateNotExist
	}
	return t, nil
}

//Init initializes all the template in the map
func (tmplts *Templates) Init() (err error) {
	for _, t := range tmplts.templates {
		err = t.Init()
		if err != nil {
			return
		}
	}
	return
}
