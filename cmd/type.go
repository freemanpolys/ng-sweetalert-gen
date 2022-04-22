package cmd

import "errors"

type HtmlFormInputType string

const (
	Text     HtmlFormInputType = "text"
	Number   HtmlFormInputType = "number"
	Textarea HtmlFormInputType = "textarea"
)

func (e HtmlFormInputType) String() string {
	return string(e)
}

var supportedType map[string]HtmlFormInputType

func init() {
	supportedType = make(map[string]HtmlFormInputType)
	supportedType["text"] = Text
	supportedType["number"] = Number
	supportedType["textarea"] = Textarea
}

func GetHtmlFormInputType(input string) (result HtmlFormInputType, err error) {
	result, ok := supportedType[input]
	if !ok {
		err = errors.New("HtmlFormInputType not found")
	}
	return
}
