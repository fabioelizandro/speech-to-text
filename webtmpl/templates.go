package webtmpl

type Template interface {
	_name() string
}

type IndexTemplate struct {
	Name string
}

func (h IndexTemplate) _name() string {
	return "index.gohtml"
}
