package webtmpl

type Template interface {
	_name() string
}

type HelloTemplate struct {
	Name string
}

func (h HelloTemplate) _name() string {
	return "hello.gohtml"
}
