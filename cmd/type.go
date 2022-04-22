package cmd

import "errors"

type FormElementType int

type FormElement struct {
	Tag   string
	Type  FormElementType
	Value string
}

const (
	Text FormElementType = iota
	Number
	Textarea
	Select
)

func (e FormElementType) String() string {
	return []string{"text", "number", "textarea", "select"}[e]
}

var supportedType map[string]FormElement

func init() {
	supportedType = make(map[string]FormElement)
	supportedType["text"] = FormElement{Tag: "input", Type: Text}
	supportedType["number"] = FormElement{Tag: "input", Type: Number}
	supportedType["textarea"] = FormElement{Tag: "textarea", Type: Textarea}
	supportedType["select"] = FormElement{Tag: "select", Type: Select}

}

func GetHtmlFormInputType(input string) (formTag FormElement, err error) {
	formTag, ok := supportedType[input]
	if !ok {
		err = errors.New("HtmlFormInputType not found")
	}
	return
}
