package cmd

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

func ProcessTmplFiles(folder string, dstFileName string, tmplFile []byte, tmplData interface{}, debug bool) {
	tmpl, err := template.New("gen").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
		"toTitle": func(i interface{}) string {
			return strings.Title(fmt.Sprintf("%v", i))
		},
		"isTextArea": func(e FormElementType) bool {
			return e == Textarea
		},
		"isSelect": func(e FormElementType) bool {
			return e == Select
		},
	}).Parse(string(tmplFile))
	tmpl = template.Must(tmpl, err)
	if err != nil {
		fmt.Print("Error Parsing template: ", err)
		os.Exit(0)
	}
	filePath := folder + "/" + dstFileName
	if debug {
		err = tmpl.Execute(os.Stderr, tmplData)
		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}
	} else {
		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			fmt.Print("Error creating file. ", err)
			os.Exit(0)
		}

		err = tmpl.Execute(file, tmplData)
		fmt.Println("create ", filePath)

	}
	if err != nil {
		fmt.Print("Error executing template. ", filePath, err)
	}

}
